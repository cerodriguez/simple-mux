package handlers

import (
    "encoding/json"
    "net/http"
    "simple-mux/internal/store"
    "strconv"
)

type ItemHandler struct {
    Store *store.Store
}

func NewItemHandler(store *store.Store) *ItemHandler {
    return &ItemHandler{Store: store}
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Name string `json:"name"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    item := h.Store.CreateItem(req.Name)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    item, exists := h.Store.GetItem(id)
    if !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var req struct {
        Name string `json:"name"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    item, exists := h.Store.UpdateItem(id, req.Name)
    if !exists {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(item)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    if idStr == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if !h.Store.DeleteItem(id) {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

