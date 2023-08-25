package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` //status can be To Do , In Progress , Done
}

var tasks []Task

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for _, task := range tasks {
		if task.ID == params["id"] {
			err := json.NewEncoder(w).Encode(task)
			if err != nil {
				fmt.Printf("%v\n", err)
				return
			}
			return
		}
	}
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

func main() {
	port := "8080"

	r := mux.NewRouter()

	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	log.Printf("Starting server at port %v", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

}
