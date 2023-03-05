package controller

import (
	"net/http"

	"github.com/bccfilkom-be/basic-jwt-2023/domain"
	"github.com/bccfilkom-be/basic-jwt-2023/middleware"
	"github.com/bccfilkom-be/basic-jwt-2023/repository"
	"github.com/bccfilkom-be/basic-jwt-2023/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	ur repository.UserRepository
}

func NewUserController(ur repository.UserRepository) *UserController {
	return &UserController{ur: ur}
}

func (uc *UserController) Register(c *gin.Context) {
	var userRegister domain.UserRegister
	if err := c.Bind(&userRegister); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	user := &domain.User{
		Username: userRegister.Username,
		Password: userRegister.Password,
		Role:     userRegister.Role,
		Name:     userRegister.Name,
	}
	if err := uc.ur.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	tokenJwt, err := middleware.GenerateToken(user.ID)
	if err != nil {
		//Handle Error
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("success", gin.H{"token": tokenJwt}))
}
