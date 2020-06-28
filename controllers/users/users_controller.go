package users

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moatazsalemVF/bookstore_user-api/domain/users"
)

//CreateUser is a func to create users
func CreateUser(c *gin.Context) {
	user := users.User{}
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	if err != nil {
		//TODO: Handle Error here
		return
	}
	c.String(http.StatusNotImplemented, "Implement me!")
}

//GetUser is a func to get users details
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
