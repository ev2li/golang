package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
	Active   bool
}

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func testRaw() {
	type Result struct {
		ID   int
		Name string
		Age  int
	}
	var result Result
	db.Raw("select id, name, age from users where name = ?", "tom").Scan(&result)
	fmt.Printf("result: %v\n", result)

	var age int
	db.Raw("select sum(age) from users").Scan(&age)
	fmt.Printf("age: %v\n", age)
}

func testRaw2() {
	db.Exec("update users set age = ? where name = ?", 100, "tom")
}

func testRaw3() {
	var user User
	db.Where("name = @myname", sql.Named("myname", "tom")).Find(&user)
	fmt.Printf("user: %v\n", user)
}

func testRaw4() {
	var name string
	var age int
	row := db.Table("users").Where("name = ?", "tom").Select("name", "age").Row()
	row.Scan(&name, &age)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)

	rows, _ := db.Model(&User{}).Where("age > ?", 18).Select("name", "age").Rows()
	for rows.Next() {
		rows.Scan(&name, &age)
		fmt.Printf("name: %v\n", name)
		fmt.Printf("age: %v\n", age)
	}
}

func main() {
	// testRaw()
	// testRaw2()
	// testRaw3()
	testRaw4()
	/* db.AutoMigrate(&User{})
	user := User{
		Name:     "kite",
		Age:      21,
		Birthday: time.Now(),
		Active:   false,
	}
	db.Create(&user) */
}
