package integration

import (
	"go-todo-app/db"
	"go-todo-app/models"
	"os"
	"testing"

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

func TestDBSelectTodos(t *testing.T) {
	db.Init()
	var todos []models.Todo
	err := db.DB.Select(&todos, "SELECT * FROM todos")
	assert.Nil(t, err)
}
