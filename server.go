package main

import (
	"user-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	user := router.Group("/api/v1/user")
	{
		uc := controllers.NewUserController()
		user.GET("/", uc.GetAll)
		user.GET("/:id", uc.Get)
	}

	router.Run(":9000")
}
