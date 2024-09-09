package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	/*
		r.GET("/get", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		r.POST("/post", func(c *gin.Context) {
			var JSON struct {
				Message  string `json:"message"`
				Location string `json:"location"`
			}

			if err := c.ShouldBindJSON(&JSON); err == nil {
				c.JSON(200, gin.H{
					"message":  JSON.Message,
					"location": JSON.Location,
				})
			} else {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			}

		})*/
	//router.SetupRouter(r)
	r.Run(":8080")
}
