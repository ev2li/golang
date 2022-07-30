package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//模型
type Product struct {
	gorm.Model //继承gorm.Model
	Code       string
	Price      uint
}

func create(db *gorm.DB) {
	//创建表
	db.AutoMigrate(&Product{})
}

func insert(db *gorm.DB) {
	//插入数据
	p := Product{
		Code:  "1001",
		Price: 100,
	}
	db.Create(&p)
}

func find(db *gorm.DB) {
	var p Product
	//查询
	// db.First(&p, 1)
	db.First(&p, "code = ?", "1001")
	fmt.Printf("p: %v\n", p)
}

func update(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	//更新单个字段
	// db.Model(&p).Update("Price", 1000)
	//更新多个字段
	// db.Model(&p).Updates(Product{Price: 200, Code: "F42"})
	db.Model(&p).Updates(map[string]interface{}{"Price": 300, "Code": "F43"})
}

func delete(db *gorm.DB) {
	var product Product
	db.First(&product, 1)
	db.Delete(&product, 1)
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// find(db)
	// update(db)
	delete(db)
}
