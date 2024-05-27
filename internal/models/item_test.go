package models

import "testing"

func TestItemCreation(t *testing.T) {
    item := Item{
        ID:   1,
        Name: "Test Item",
    }

    if item.ID != 1 {
        t.Errorf("expected ID to be 1, got %d", item.ID)
    }

    if item.Name != "Test Item" {
        t.Errorf("expected Name to be 'Test Item', got '%s'", item.Name)
    }
}

