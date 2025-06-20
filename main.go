package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-todo-app/db"
	"go-todo-app/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()

	db.Init()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ToDo API is running"))
	})

	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	handler := cors.AllowAll().Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
