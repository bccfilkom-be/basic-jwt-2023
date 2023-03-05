package controller

import (
	"net/http"

	"github.com/bccfilkom-be/basic-jwt-2023/domain"
	"github.com/bccfilkom-be/basic-jwt-2023/repository"
	"github.com/bccfilkom-be/basic-jwt-2023/utils"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) *TodoController {
	return &TodoController{tr: tr}
}

func (tc *TodoController) Create(c *gin.Context) {
	var inputTodo domain.TodoRequest
	if err := c.Bind(&inputTodo); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}

	idUser := c.MustGet("id").(uint)

	// Fungsi buat query pengguna where id = id

	todo := &domain.Todo{
		Activity: inputTodo.Activity,
		IdUser:   idUser,
	}

	if err := tc.tr.Create(todo); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success", todo))
}
