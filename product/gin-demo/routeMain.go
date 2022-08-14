package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("/hello", func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath())
		name := ctx.Query("name")
		fmt.Println(name)
		ctx.Writer.Write([]byte("success"))
	})

	e.POST("/login", func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath())
		name, exists := ctx.GetPostForm("username")
		if exists {
			fmt.Println(name)
		}
		ctx.Writer.Write([]byte("success"))
	})

	if err := e.Run(":8080"); err != nil {
		recover()
	}

}
