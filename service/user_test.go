package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	array := []string{"user", "todo2", "todo3"}

	userService := &UserService{}

	err := userService.AddTodos(array[:])

	assert.Nil(t, err)
}
