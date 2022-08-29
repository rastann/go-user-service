package main

import (
	"tutorial/user-service/database"
	"tutorial/user-service/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetDB()

	userRouters := routers.GetUserRoute(db)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", userRouters.GetUsers)
	}

	router.Run("localhost:8080")
}
