package app

import (
	"github.com/moatazsalemVF/bookstore_user-api/controllers/ping"
	"github.com/moatazsalemVF/bookstore_user-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
