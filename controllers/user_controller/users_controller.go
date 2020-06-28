package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser is a func to create users
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

//GetUser is a func to get users details
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
