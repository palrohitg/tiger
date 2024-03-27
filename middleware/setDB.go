package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const DB = "DB"

var DBC *gorm.DB
var DbcConsumer *gorm.DB

func SetDBToContext(db *gorm.DB) gin.HandlerFunc {
	DBC = db

	return func(c *gin.Context) {
		c.Set(DB, db)
		c.Next()
	}
}
