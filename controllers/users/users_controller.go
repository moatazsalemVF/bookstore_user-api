package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moatazsalemVF/bookstore_user-api/domain/users"
	"github.com/moatazsalemVF/bookstore_user-api/services"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
)

//CreateUser is a func to create users
func CreateUser(c *gin.Context) {
	user := users.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

//GetUser is a func to get users details
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
