package handler

import "github.com/gin-gonic/gin"

func GetHelloMessage() string {
	return "Hello, dari Gin"
}

func RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": GetHelloMessage(),
	})
}

func PostHandler(c *gin.Context) {
	var JSON struct {
		Message string `json:"message"`
	}

	if err := c.ShouldBindJSON(&JSON); err == nil {
		c.JSON(200, gin.H{
			"message": JSON.Message,
		})
	} else {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
}
