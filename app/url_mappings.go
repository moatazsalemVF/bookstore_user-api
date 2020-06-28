package app

import "github.com/moatazsalemVF/bookstore_user-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
