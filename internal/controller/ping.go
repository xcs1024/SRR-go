package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(content *gin.Context) {
	content.String(http.StatusOK, "pong")
}
