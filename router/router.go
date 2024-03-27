package router

import (
	"github.com/gin-gonic/gin"
	handlers "tigerhall/apps/tiger-handler/handler"
	"tigerhall/apps/tiger/urls"
	"tigerhall/dtos"
)

func Initialize(r *gin.Engine) {
	println("Initializing router")
	r.GET("/healthcheck/", handlers.Healthcheck)
	setupApiRoutes(r)
	println("Router Initialized")
}

func setupApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	addV1URLs(dtos.RouterGroups{
		NoAuth: api,
	})
}

func addV1URLs(groups dtos.RouterGroups) {
	v1 := groups.NoAuth.Group("/v1")
	routerGroups := dtos.RouterGroups{NoAuth: v1}
	urls.AddV1URLs(routerGroups)
}
