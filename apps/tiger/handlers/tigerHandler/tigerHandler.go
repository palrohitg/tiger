package tigerHandler

import (
	"github.com/gin-gonic/gin"
	"tigerhall/apps/tiger/controllers/tigerController"
	"tigerhall/config"
	"tigerhall/dtos"
)

func CreateUserHandler(ctx *gin.Context) {
	/*
		todo: Password hashing algorithm should be use here
	*/
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
	response := map[string]interface{}{
		"user_info": data,
	}
	config.SendSuccessResponse(ctx, config.SuccessStatus, config.SuccessMsg, response)
}

func LoginHandler(ctx *gin.Context) {
	var login dtos.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	data, err := tigerController.LoginUser(login)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	response := map[string]interface{}{
		"token_info": data,
	}
	config.SendSuccessResponse(ctx, config.SuccessStatus, config.SuccessMsg, response)
}

func CreateTigerHandler(ctx *gin.Context) {
	var tigerDetails dtos.TigerDetails
	err := ctx.ShouldBindJSON(&tigerDetails)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	_, err = tigerController.CreateTiger(tigerDetails)
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	data := make(map[string]interface{})
	config.SendSuccessResponse(ctx, config.SuccessStatus, config.SuccessMsg, data)

}

func GetTigerListHandler(ctx *gin.Context) {
	// todo: list of tigers tiger details and their last information
	// Also add the pagination support here
	data, err := tigerController.TigerDetails()
	if err != nil {
		config.SendErrorResponse(ctx, config.ErrStatus, err.Error())
		return
	}
	config.SendSuccessResponse(ctx, config.SuccessStatus, config.SuccessMsg, data)
}

func GetTigerSightHandler(ctx *gin.Context) {

}

func CreateSightTigerHandler(ctx *gin.Context) {
	/*
		If the last lat-long distance after 5km then it should
		created, check from the previous.
	*/
	//var tigerSightDetails dtos.TigerSightDetails
}
