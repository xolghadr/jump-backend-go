package main

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn     *websocket.Conn
	username string
}

var clients = make(map[string][]*Client)
var mu sync.Mutex

func wsChatRoom(c echo.Context) error {
	roomId := c.Param("roomId")
	username := c.Param("username")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := &Client{conn: conn, username: username}

	mu.Lock()
	clients[roomId] = append(clients[roomId], client)
	mu.Unlock()

	defer func() {
		mu.Lock()
		for i, c := range clients[roomId] {
			if c == client {
				clients[roomId] = append(clients[roomId][:i], clients[roomId][i+1:]...)
				break
			}
		}
		mu.Unlock()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		broadcastMessage(roomId, username, string(msg))
	}

	return nil
}

func broadcastMessage(roomId, username, message string) {
	mu.Lock()
	defer mu.Unlock()

	for _, client := range clients[roomId] {
		if client.username != username {
			err := client.conn.WriteMessage(websocket.TextMessage, []byte(username+": "+message))
			if err != nil {
				client.conn.Close()
				// Remove the client from the list if there's an error
				removeClient(roomId, client)
			}
		}
	}
}

func removeClient(roomId string, client *Client) {
	mu.Lock()
	defer mu.Unlock()

	for i, c := range clients[roomId] {
		if c == client {
			clients[roomId] = append(clients[roomId][:i], clients[roomId][i+1:]...)
			break
		}
	}
}

func main() {
	e := echo.New()
	e.GET("/ws/chat/:roomId/user/:username", wsChatRoom)
	e.Logger.Fatal(e.Start(":8080"))
}
