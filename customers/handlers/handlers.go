package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Greet(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World!")
	return
}

func Health(msg string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, msg)
	}
}
