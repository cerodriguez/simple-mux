package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "simple-mux/internal/models"
    "simple-mux/internal/store"
    "strconv"
    "testing"
)

func TestCreateItemHandler(t *testing.T) {
    s := store.NewStore()
    h := NewItemHandler(s)

    rr := httptest.NewRecorder()
    body := []byte(`{"name":"Test Item"}`)
    req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    http.HandlerFunc(h.CreateItem).ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var item models.Item
    if err := json.NewDecoder(rr.Body).Decode(&item); err != nil {
        t.Fatal(err)
    }

    if item.Name != "Test Item" {
        t.Errorf("expected Name to be 'Test Item', got '%s'", item.Name)
    }
}

func TestGetItemHandler(t *testing.T) {
    s := store.NewStore()
    h := NewItemHandler(s)
    createdItem := s.CreateItem("Test Item")

    rr := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "/item?id="+strconv.Itoa(createdItem.ID), nil)
    if err != nil {
        t.Fatal(err)
    }

    http.HandlerFunc(h.GetItem).ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var item models.Item
    if err := json.NewDecoder(rr.Body).Decode(&item); err != nil {
        t.Fatal(err)
    }

    if item != createdItem {
        t.Errorf("expected item to be %v, got %v", createdItem, item)
    }
}

func TestUpdateItemHandler(t *testing.T) {
    s := store.NewStore()
    h := NewItemHandler(s)
    createdItem := s.CreateItem("Test Item")

    rr := httptest.NewRecorder()
    body := []byte(`{"name":"Updated Item"}`)
    req, err := http.NewRequest("PUT", "/item?id="+strconv.Itoa(createdItem.ID), bytes.NewBuffer(body))
    if err != nil {
        t.Fatal(err)
    }

    http.HandlerFunc(h.UpdateItem).ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var item models.Item
    if err := json.NewDecoder(rr.Body).Decode(&item); err != nil {
        t.Fatal(err)
    }

    if item.Name != "Updated Item" {
        t.Errorf("expected Name to be 'Updated Item', got '%s'", item.Name)
    }
}

func TestDeleteItemHandler(t *testing.T) {
    s := store.NewStore()
    h := NewItemHandler(s)
    createdItem := s.CreateItem("Test Item")

    rr := httptest.NewRecorder()
    req, err := http.NewRequest("DELETE", "/item?id="+strconv.Itoa(createdItem.ID), nil)
    if err != nil {
        t.Fatal(err)
    }

    http.HandlerFunc(h.DeleteItem).ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNoContent {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusNoContent)
    }
}

