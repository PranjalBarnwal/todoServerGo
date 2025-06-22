package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-todo-app/db"
	"go-todo-app/handlers"
	"go-todo-app/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Failed to load .env: " + err.Error())
	}
	os.Exit(m.Run())
}

func setupRouter() *mux.Router {
	db.Init()
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	return r
}

func TestCreateTodoAPI(t *testing.T) {
	router := setupRouter()
	todo := models.Todo{Title: "Test Task", Completed: false}
	body, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetTodosAPI(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/todos", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateTodoAPI(t *testing.T) {
	router := setupRouter()

	initial := models.Todo{Title: "Original Title", Completed: false}
	body, _ := json.Marshal(initial)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var created models.Todo
	_ = json.NewDecoder(resp.Body).Decode(&created)

	updated := models.Todo{Title: "Updated Title", Completed: true}
	updateBody, _ := json.Marshal(updated)
	updateReq, _ := http.NewRequest("PUT", "/todos/"+fmt.Sprint(created.ID), bytes.NewBuffer(updateBody))
	updateReq.Header.Set("Content-Type", "application/json")
	updateResp := httptest.NewRecorder()
	router.ServeHTTP(updateResp, updateReq)

	assert.Equal(t, http.StatusNoContent, updateResp.Code)
}

func TestDeleteTodoAPI(t *testing.T) {
	router := setupRouter()

	todo := models.Todo{Title: "To be deleted", Completed: false}
	body, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var created models.Todo
	_ = json.NewDecoder(resp.Body).Decode(&created)

	deleteReq, _ := http.NewRequest("DELETE", "/todos/"+fmt.Sprint(created.ID), nil)
	deleteResp := httptest.NewRecorder()
	router.ServeHTTP(deleteResp, deleteReq)

	assert.Equal(t, http.StatusNoContent, deleteResp.Code)
}
