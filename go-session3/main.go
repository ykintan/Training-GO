package main

import (
	"training-go/go-session3/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// empty main function
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(r)

	r.Run("localhost:8080")
}
