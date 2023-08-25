package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` //status can be To Do , In Progress , Done
}

var tasks []Task

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

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

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

}
