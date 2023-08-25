package main

import (
	"github.com/gorilla/mux"
	"log"
)

var Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` //status can be To Do , In Progress , Done
}

func main() {
	port := "8080"

	r := mux.NewRouter()

	r.HandleFunc("/tasks", getTasks)
	r.HandleFunc("/tasks/{id}", getTask)
	r.HandleFunc("/tasks", createTask)
	r.HandleFunc("/tasks/{id}", deleteTask)
	r.HandleFunc("/tasks/{id}", updateTask)

	log.Printf("Starting server at port %v", port)

}
