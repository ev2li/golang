package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	/* fmt.Println("hello world")
	service.TestUserService()
	service.TestCustomerService() */
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
