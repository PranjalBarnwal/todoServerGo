package unit

import (
	"encoding/json"
	"go-todo-app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoStruct(t *testing.T) {
	todo := models.Todo{
		ID:        1,
		Title:     "Read book",
		Completed: false,
	}
	assert.Equal(t, 1, todo.ID)
	assert.Equal(t, "Read book", todo.Title)
	assert.False(t, todo.Completed)
}

func TestTodoJSONMarshalling(t *testing.T) {
	todo := models.Todo{ID: 1, Title: "Read book", Completed: false}
	data, err := json.Marshal(todo)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "Read book")
}
