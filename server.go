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
		user.POST("/", uc.Create)
		user.PUT("/", uc.Update)
		user.PATCH("/status", uc.UpdateStatus)
		user.PATCH("/password", uc.UpdatePassword)
		user.DELETE("/:id", uc.Delete)
	}

	router.Run(":9000")
}
