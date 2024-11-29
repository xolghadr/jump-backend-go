package main

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

type Store struct {
	Products map[string]*Product
	mux      sync.Mutex
}

type Product struct {
	name       string
	stockCount int
	price      float64
}

func NewStore() *Store {
	newStore := Store{
		Products: make(map[string]*Product, 0),
	}

	return &newStore
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, ok := s.Products[strings.ToLower(name)]
	if ok {
		return fmt.Errorf("%s already exists", name)
	}

	if price <= 0 {
		return fmt.Errorf("price should be positive")
	}

	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}

	newProduct := &Product{
		name:       name,
		price:      price,
		stockCount: count,
	}

	s.Products[strings.ToLower(name)] = newProduct
	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	name = strings.ToLower(name)
	p, ok := s.Products[name]

	if !ok {
		return 0, fmt.Errorf("invalid product name")
	}

	return p.stockCount, nil
}

func (s *Store) GetProductPrice(name string) (float64, error) {
	name = strings.ToLower(name)
	p, ok := s.Products[name]

	if !ok {
		return 0, fmt.Errorf("invalid product name")
	}

	return p.price, nil
}

func (s *Store) Order(name string, count int) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}
	p, ok := s.Products[strings.ToLower(name)]
	if !ok {
		return fmt.Errorf("invalid product name")
	}
	if p.stockCount == 0 {
		return fmt.Errorf("there is no %s in the store", name)
	}
	if p.stockCount < count {
		return fmt.Errorf("not enough %s in the store. there are %d left", name, p.stockCount)
	}

	p.stockCount -= count
	return nil
}

func (s *Store) ProductsList() ([]string, error) {
	products := make([]string, 0)
	for _, p := range s.Products {
		if p.stockCount > 0 {
			products = append(products, p.name)
		}
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("store is empty")
	}

	slices.Sort(products)
	return products, nil
}
