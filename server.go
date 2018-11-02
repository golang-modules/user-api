package main

import (
	"user-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	//"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//store := cookie.NewStore([]byte("secret"))
	store := memstore.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mysession", store))
	router.Use(cors.Default())

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

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	router.Run(":9000")
}
