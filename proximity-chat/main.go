package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type Player struct {
	channel    chan string
	name       string
	currentMap Map
}

type Map struct {
	channel chan string
	id      int
	players map[string]*Player
}

type Game struct {
	maps    map[int]Map
	players map[string]Player
	mux     *sync.Mutex
}

func NewGame(mapIds []int) (*Game, error) {
	maps := make(map[int]Map, len(mapIds))
	for _, id := range mapIds {
		if id <= 0 {
			return nil, errors.New("map id out of range")
		}
		maps[id] = Map{make(chan string), id, make(map[string]*Player)}
	}
	newGame := Game{
		maps:    maps,
		mux:     &sync.Mutex{},
		players: make(map[string]Player),
	}
	return &newGame, nil
}

func (g *Game) ConnectPlayer(name string) error {
	defer g.mux.Unlock()
	g.mux.Lock()
	if _, ok := g.players[strings.ToLower(name)]; ok {
		return errors.New("player already exists")
	}

	g.players[strings.ToLower(name)] = Player{name: name, channel: make(chan string)}
	return nil
}

func (g *Game) SwitchPlayerMap(name string, mapId int) error {
	newMap, ok := g.maps[mapId]
	if !ok {
		return errors.New("map not found")
	}
	player, ok := g.players[strings.ToLower(name)]
	if !ok {
		return errors.New("player not found")
	}
	if player.currentMap.id == mapId {
		return errors.New("already connected")
	}
	delete(player.currentMap.players, player.GetLowerName())
	player.currentMap = newMap
	newMap.players[player.GetLowerName()] = &player
	return nil
}

func (g *Game) GetPlayer(name string) (*Player, error) {
	player, ok := g.players[strings.ToLower(name)]
	if !ok {
		return nil, errors.New("player not found")
	}
	return &player, nil
}

func (g *Game) GetMap(mapId int) (*Map, error) {
	m, ok := g.maps[mapId]
	if !ok {
		return nil, errors.New("map not found")
	}
	return &m, nil
}

func (m *Map) FanOutMessages() {
	for _, player := range m.players {
		player.channel <- <-m.channel
	}
}

func (p *Player) GetChannel() <-chan string {
	return p.channel
}

func (p *Player) GetLowerName() string {
	return strings.ToLower(p.name)
}

func ToPascalCase(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func (p *Player) SendMessage(msg string) error {
	if msg == "" {
		return errors.New("message is empty")
	}
	p.currentMap.channel <- fmt.Sprintf("%s says: %s", ToPascalCase(p.name), msg)
	return nil
}

func (p *Player) GetName() string {
	return p.name
}
