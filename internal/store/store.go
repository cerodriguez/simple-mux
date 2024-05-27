package store

import (
    "simple-mux/internal/models"
    "sync"
)

type Store struct {
    items  map[int]models.Item
    nextID int
    mu     sync.Mutex
}

func NewStore() *Store {
    return &Store{
        items:  make(map[int]models.Item),
        nextID: 1,
    }
}

func (s *Store) CreateItem(name string) models.Item {
    s.mu.Lock()
    defer s.mu.Unlock()

    item := models.Item{
        ID:   s.nextID,
        Name: name,
    }
    s.items[s.nextID] = item
    s.nextID++
    return item
}

func (s *Store) GetItem(id int) (models.Item, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()

    item, exists := s.items[id]
    return item, exists
}

func (s *Store) UpdateItem(id int, name string) (models.Item, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()

    item, exists := s.items[id]
    if !exists {
        return models.Item{}, false
    }
    item.Name = name
    s.items[id] = item
    return item, true
}

func (s *Store) DeleteItem(id int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()

    _, exists := s.items[id]
    if !exists {
        return false
    }
    delete(s.items, id)
    return true
}

