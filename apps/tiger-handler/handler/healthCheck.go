package handlers

import (
	"net/http"
	"tigerhall/config"
	"time"

	"github.com/gin-gonic/gin"
)

var t = time.Now()

var rounded = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

var currentTime = rounded.Add(time.Hour * 4)

func Healthcheck(ctx *gin.Context) {
	response := map[string]interface{}{
		"message": config.SuccessMsg,
		"status":  config.SuccessStatus,
		"msg":     currentTime,
	}
	ctx.JSON(http.StatusOK, response)
}
