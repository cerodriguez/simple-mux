package store

import (
    "testing"
)

func TestCreateItem(t *testing.T) {
    s := NewStore()
    item := s.CreateItem("Test Item")

    if item.ID != 1 {
        t.Errorf("expected ID to be 1, got %d", item.ID)
    }

    if item.Name != "Test Item" {
        t.Errorf("expected Name to be 'Test Item', got '%s'", item.Name)
    }
}

func TestGetItem(t *testing.T) {
    s := NewStore()
    createdItem := s.CreateItem("Test Item")
    item, exists := s.GetItem(createdItem.ID)

    if !exists {
        t.Fatalf("expected item to exist")
    }

    if item != createdItem {
        t.Errorf("expected item to be %v, got %v", createdItem, item)
    }
}

func TestUpdateItem(t *testing.T) {
    s := NewStore()
    createdItem := s.CreateItem("Test Item")
    updatedItem, exists := s.UpdateItem(createdItem.ID, "Updated Item")

    if !exists {
        t.Fatalf("expected item to exist")
    }

    if updatedItem.Name != "Updated Item" {
        t.Errorf("expected Name to be 'Updated Item', got '%s'", updatedItem.Name)
    }
}

func TestDeleteItem(t *testing.T) {
    s := NewStore()
    createdItem := s.CreateItem("Test Item")
    deleted := s.DeleteItem(createdItem.ID)

    if !deleted {
        t.Fatalf("expected item to be deleted")
    }

    _, exists := s.GetItem(createdItem.ID)
    if exists {
        t.Errorf("expected item to not exist")
    }
}

