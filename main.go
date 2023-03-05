package main

import (
	"fmt"
	"os"

	"github.com/bccfilkom-be/basic-jwt-2023/controller"
	"github.com/bccfilkom-be/basic-jwt-2023/middleware"
	"github.com/bccfilkom-be/basic-jwt-2023/repository"
	"github.com/bccfilkom-be/basic-jwt-2023/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		//Handle Error
		panic(err.Error())
	}
	dbDriver, err := utils.ReadEnvDatabase()
	if err != nil {
		//Handle Error
		panic(err.Error())
	}
	db, err := utils.MakeConnection(dbDriver)
	if err != nil {
		//Handle Error
		panic(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	todoRepository := repository.NewTodoRepository(db)
	todoController := controller.NewTodoController(todoRepository)

	r := gin.New()

	r.POST("user/register", userController.Register)
	r.POST("todo", middleware.ValidateToken(), todoController.Create)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
