package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
				log.Printf("%v\n", err)
				return
			}
			return
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, task := range tasks {
		if task.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)

			err := json.NewEncoder(w).Encode(tasks)
			if err != nil {
				log.Printf("%v\n", err)
				return
			}
			return
		}
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var newTask Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)

	newTask.ID = strconv.Itoa(rand.Intn(9999999999))
	tasks = append(tasks, newTask)

	err := json.NewEncoder(w).Encode(newTask)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
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
