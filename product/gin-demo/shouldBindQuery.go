package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type User struct {
	Username string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}

func main() {
	e := gin.Default()
	e.GET("/hello", func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath())
		var student Student
		err := ctx.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(student.Name, student.Classes)
		ctx.Writer.Write([]byte("success"))
	})

	e.POST("/login", func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath())
		var user User
		err := ctx.ShouldBind(&user)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(user.Username, user.Phone, user.Password)
		ctx.Writer.Write([]byte("success"))
	})

	e.POST("/register", func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath())
		var person Person
		err := ctx.BindJSON(&person)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(person.Name, person.Sex, person.Age)
		ctx.Writer.Write([]byte("success"))
	})
	e.Run()
}
