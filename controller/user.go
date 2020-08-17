package controller

import (

	// "github.com/kzpolicy/amb-todo/service"

	"local.packages/models"
	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	User models.User `json:"user"`
}

// func GetTodo(c *gin.Context) {

// 	userService := service.userService{}
// 	TodoList := userService.GetTodoList()

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "ok",
// 		"data":    {TodoList},
// 	})
// }

func CreateUser(c *gin.Context) {

	var userForm UserForm
	c.BindJSON(&userForm)

	userService := service.UserService{}
	err := userService.AddTodos(userForm.User)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
