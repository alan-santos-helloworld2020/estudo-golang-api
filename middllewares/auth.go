package middllewares

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func AuthUser() gin.HandlerFunc {

     err := godotenv.Load(".env")
	 if err != nil {
        log.Fatal("Erro ao ler variaveis de ambiente")
    }

	bearerToken := os.Getenv("BEARER_TOKEN")



	return func(ctx *gin.Context) {

		fmt.Println(bearerToken)

		ctx.Next()

	}
}