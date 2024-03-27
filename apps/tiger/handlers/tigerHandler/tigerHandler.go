package tigerHandler

import (
	"github.com/gin-gonic/gin"
	"tigerhall/apps/tiger/controllers/tigerController"
	"tigerhall/config"
	"tigerhall/dtos"
)

func CreateUserHandler(ctx *gin.Context) {
	var user dtos.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	data, err := tigerController.CreateUser(user)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	config.SendSuccessResponse(ctx, config.SuccessStatus, config.SuccessMsg, data)
}

func LoginHandler(ctx *gin.Context) {

}

func CreateTigerHandler(ctx *gin.Context) {

}

func GetTigerSightHandler(ctx *gin.Context) {

}
