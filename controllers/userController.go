package controllers

import (
	"app/app/database"
	"app/app/models"

	"github.com/gin-gonic/gin"
)

type UserUri struct {
	Id int `uri:"id"`
}

func GetAlUsers (ctx *gin.Context) {

	users := database.GetUsers()
	var _ = ctx.BindJSON(&users)
	ctx.JSON(200, users)
}

func GetUserById(ctx *gin.Context) {

	var u UserUri
	if err := ctx.ShouldBindUri(&u); err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	
	var us models.User
	for _, v := range database.GetUsers() {

		if v.Id == u.Id {
			us = v
		}
		
	}

	ctx.JSON(200, us)
}