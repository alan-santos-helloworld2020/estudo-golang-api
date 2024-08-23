package main

import (
	"app/app/controllers"
	"app/app/middllewares"

	"github.com/gin-gonic/gin"
)



func main() {

	server := gin.Default()	
	
	usuario := server.Group("/usuario")
	usuario.Use()
	{
		usuario.GET("/", controllers.GetAlUsers)
		usuario.GET("/:id", controllers.GetUserById)
	}
	
	
	cliente := server.Group("/cliente")
	cliente.Use(middllewares.AuthUser())
	{
		cliente.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{

			})
		})

	}

	server.Run()

}
