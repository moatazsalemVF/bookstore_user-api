package app

import (
	"github.com/moatazsalemVF/bookstore_user-api/controllers/ping_controller"
	"github.com/moatazsalemVF/bookstore_user-api/controllers/user_controller"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.GET("/users/:user_id", user_controller.GetUser)
	router.POST("/users", user_controller.CreateUser)
}
