package main

import (
	"gouserapi/configs"
	"gouserapi/controllers"
	"gouserapi/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitDB()
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)

	r.Run(":8080")
}
