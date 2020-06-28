package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

//StartApplication is a function to act
//as an entry point for the app
func StartApplication() {
	mapUrls()

	router.Run(":8080")
}
