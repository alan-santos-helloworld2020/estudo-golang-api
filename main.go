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
		cliente.GET("/", controllers.GetClientes)
		cliente.GET("/:id", controllers.GetClienteById)
		cliente.POST("/", controllers.PostSalvarCliente)

	}

	server.Run()

}
