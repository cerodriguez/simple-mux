package main

import (
    "fmt"
    "net/http"
    "simple-mux/internal/handlers"
    "simple-mux/internal/store"
)

func main() {
    store := store.NewStore()
    itemHandler := handlers.NewItemHandler(store)

    mux := http.NewServeMux()
    mux.HandleFunc("POST /items", itemHandler.CreateItem)
    mux.HandleFunc("GET /item", itemHandler.GetItem)
    mux.HandleFunc("PUT /item", itemHandler.UpdateItem)
    mux.HandleFunc("DELETE /item", itemHandler.DeleteItem)

    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", mux)
}

