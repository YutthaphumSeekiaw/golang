package handlers

import (
    "encoding/json"
    "net/http"
    "sync"

    "GoSwaggerApp/internal/models"
)

var (
    items  = []models.Item{}
    nextID = 1
    mu     sync.Mutex
)

func GetItems(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    item.ID = nextID
    nextID++
    items = append(items, item)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request, id int) {
    mu.Lock()
    defer mu.Unlock()

    for _, item := range items {
        if item.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }

    http.Error(w, "Item not found", http.StatusNotFound)
}

func UpdateItem(w http.ResponseWriter, r *http.Request, id int) {
    mu.Lock()
    defer mu.Unlock()

    for i, item := range items {
        if item.ID == id {
            var updatedItem models.Item
            if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                return
            }

            updatedItem.ID = id
            items[i] = updatedItem

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }

    http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItem(w http.ResponseWriter, r *http.Request, id int) {
    mu.Lock()
    defer mu.Unlock()

    for i, item := range items {
        if item.ID == id {
            items = append(items[:i], items[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "Item not found", http.StatusNotFound)
}