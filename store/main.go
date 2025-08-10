package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Product struct {
	Name string
	Price float64
	Count int
}

type Store struct {
	products map[string]Product
	mu sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		products: make(map[string]Product),
	}
}

func (s *Store) AddProduct(name string, price float64, count int) error {
	lowerName := strings.ToLower(name)
	if price <= 0 {
		return fmt.Errorf("price should be positive")
	}
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _,exist := s.products[lowerName]; exist {
		return fmt.Errorf("%s already exists", name)
	}
	s.products[lowerName] = Product{Name: name, Price: price, Count: count}

	return nil
}

func (s *Store) GetProductCount(name string) (int, error) {
	lowerName := strings.ToLower(name)
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exist := s.products[lowerName]
	if !exist {
		return 0, fmt.Errorf("invalid product name")
	}

	return s.products[lowerName].Count, nil
}

func (s *Store) GetProductPrice(name string) (float64, error) {
	lowerName := strings.ToLower(name)
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exist := s.products[lowerName]
	if !exist {
		return 0, fmt.Errorf("invalid product name")
	}
	return s.products[lowerName].Price, nil
}

func (s *Store) Order(name string, count int) error {
	lowerName := strings.ToLower(name)
	s.mu.Lock()
	defer s.mu.Unlock()
	if count <= 0 {
		return fmt.Errorf("count should be positive")
	}
	
	_, exist := s.products[lowerName]
	if !exist {
		return fmt.Errorf("invalid product name")
	}
	if s.products[lowerName].Count == 0 {
		return fmt.Errorf("there is no %s in the store", name)
	}
	if s.products[lowerName].Count < count {
		return fmt.Errorf("not enough %s in the store. there are %d left", name, s.products[lowerName].Count)
	}
	product := s.products[lowerName]
	product.Count -= count
	s.products[lowerName] = product

	return nil
}

func (s *Store) ProductsList() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	products := make([]string, 0, len(s.products))
	for _,product := range s.products {
		if product.Count > 0 {
			products = append(products, product.Name)
		}
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("store is empty")
	}
	sort.Strings(products)
	return products, nil
}