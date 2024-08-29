package controllers

import (
	"app/app/database"
	"app/app/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

 type ClienteId struct{
	Id int `uri:"id"`
 } 


func GetClientes (ctx *gin.Context) {

	res :=  database.GetClientes();
	 ctx.JSON(200, res)
}


func GetClienteById(ctx *gin.Context) {

	var c *ClienteId
	if err := ctx.ShouldBindUri(&c); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	res :=  database.GetClienteById(&c.Id);
	 ctx.JSON(200, res)
}



func PostSalvarCliente (ctx *gin.Context) {

	var cl models.ClienteDto
    if err := ctx.ShouldBindJSON(&cl); err != nil {
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }


	fmt.Println(&cl)
	database.SalvarCliente(&cl)    
	ctx.JSON(201, gin.H{"message":"cliente salvo com sucesso:)"})
}