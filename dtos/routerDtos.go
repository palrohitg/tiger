package dtos

import "github.com/gin-gonic/gin"

type RouterGroups struct {
	NoAuth     *gin.RouterGroup
	JwtAuth    *gin.RouterGroup
	ApiKeyAuth *gin.RouterGroup
}
