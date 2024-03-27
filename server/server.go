package server

import (
	"github.com/gin-gonic/gin"
	requestId "github.com/thanhhh/gin-requestid"
	"tigerhall/db"
	"tigerhall/middleware"
	"tigerhall/router"
)

func Setup() *gin.Engine {
	routerR := gin.Default()
	database := db.Connect()
	routerR.Use(requestId.RequestID())
	routerR.Use(middleware.SetDBToContext(database))
	router.Initialize(routerR)
	return routerR
}
