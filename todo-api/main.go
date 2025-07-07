package main

import (
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	Title string `json: "title"`
	Done  bool   `json: "done"`
}

var Tasks = []Task{
	{Title: "Học Go", Done: false},
	{Title: "Viet APi", Done: true},
}



func main() {
	RegisterRoutes()
	fmt.Println("🚀 Server chạy tại http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}