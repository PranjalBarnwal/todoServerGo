package handlers

import (
	"encoding/json"
	"go-todo-app/db"
	"go-todo-app/models"
	"net/http"

	"github.com/gorilla/mux"
)

// Get All Todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// Create Todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	err := db.DB.QueryRow(
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id",
		todo.Title, todo.Completed,
	).Scan(&todo.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// Update Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	_, err := db.DB.Exec(
		"UPDATE todos SET title=$1, completed=$2 WHERE id=$3",
		todo.Title, todo.Completed, id,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Delete Todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := db.DB.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
