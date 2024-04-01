package urls

import (
	"tigerhall/apps/tiger/handlers/tigerHandler"
	"tigerhall/dtos"
)

/*
1. Create User
2. Create Tiger
3. Listing Tiger based on LastSees
4. Tiger History Need to been shown here
5. Email and Password
*/
func AddV1URLs(groups dtos.RouterGroups) {
	noAuthRg := groups.NoAuth
	TigerGroup := noAuthRg.Group("/tiger-service")
	TigerGroup.POST("/create-user", tigerHandler.CreateUserHandler)
	TigerGroup.POST("/login", tigerHandler.LoginHandler)
	TigerGroup.POST("/create-tiger", tigerHandler.CreateTigerHandler)
	TigerGroup.GET("/list", tigerHandler.GetTigerListHandler)
	TigerGroup.GET("/tiger-sight", tigerHandler.GetTigerSightHandler)
	TigerGroup.POST("/create-sight", tigerHandler.CreateSightTigerHandler)
}
