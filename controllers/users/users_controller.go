package users

import (
	"net/http"
	"strconv"

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

//UpdateUser is a func to update users
func UpdateUser(c *gin.Context) {
	user := users.User{}

	idstr := c.Params.ByName("user_id")
	id, errparse := strconv.ParseInt(idstr, 10, 64)

	if errparse != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid Id"))
		return
	}

	if id == 0 {
		restErr := errors.NewBadRequestError("Missing User ID")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = id

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	result, saveErr := services.UpdateUser(user, isPartial)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

//GetUser is a func to get users details
func GetUser(c *gin.Context) {
	idstr := c.Params.ByName("user_id")
	id, errparse := strconv.ParseInt(idstr, 10, 64)

	if errparse != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid Id"))
		return
	}
	result, err := services.GetUser(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

//DeleteUser is a func to update users
func DeleteUser(c *gin.Context) {
	idstr := c.Params.ByName("user_id")
	id, errparse := strconv.ParseInt(idstr, 10, 64)

	if errparse != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("Invalid Id"))
		return
	}

	saveErr := services.DeleteUser(id)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	user := users.User{}
	user.ID = id

	c.JSON(http.StatusOK, user)
}
