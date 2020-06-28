package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping is a function to check MS health
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
