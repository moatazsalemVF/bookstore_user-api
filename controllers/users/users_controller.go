package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moatazsalemVF/bookstore_user-api/domain/users"
)

//CreateUser is a func to create users
func CreateUser(c *gin.Context) {
	user := users.User{}
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	if err != nil {
		//TODO: Handle Error here
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle JSON Error here
		return
	}
	fmt.Println(user)
	c.String(http.StatusNotImplemented, "Implement me!")
}

//GetUser is a func to get users details
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
