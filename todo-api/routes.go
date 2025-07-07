package main

import "net/http"

func RegisterRoutes() {
    http.HandleFunc("/todos", handleTodos)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        GetTodos(w, r)
    case http.MethodPost:
        CreateTodo(w, r)
    default:
        http.Error(w, "🚫 Phương thức không hỗ trợ", http.StatusMethodNotAllowed)
    }
}
