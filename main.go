package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

var Todos []Todo

func ErrorHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", allTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", getTodoById).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	// Init Todos slice
	Todos = append(Todos, Todo{Id: 1, Todo: "Aprender Go", Done: false})

	log.Fatal(http.ListenAndServe(":3001", r))
}

func allTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	ErrorHandle(err)

	todo.Id = len(Todos) + 1
	todo.Done = false
	Todos = append(Todos, todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range Todos {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Todo{})
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for index, item := range Todos {
		if item.Id == id {
			Todos = append(Todos[:index], Todos[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.Id = id
			Todos = append(Todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	ErrorHandle(err)

	for index, item := range Todos {
		if item.Id == id {
			Todos = append(Todos[:index], Todos[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
