package main

import (
	"training-go/go-session2/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//router.SetupRouter(r)
	router.SetupRouter(r)

	r.Run("localhost:8080")
}
