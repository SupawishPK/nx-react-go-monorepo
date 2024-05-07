package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type todo struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	addTodo := todo{
		Name:      "Create a mono repo",
		Completed: false,
	}
	json.NewEncoder(w).Encode(addTodo)
}

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API is up and running")
}

func handleRequest() {
	http.HandleFunc("/", health)
	http.HandleFunc("/todos", getTodos)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
