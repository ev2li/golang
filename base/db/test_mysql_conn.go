package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertData() {
	sqlStr := "insert into user_tbl(username, password) values(?,?)"
	rt, err := db.Exec(sqlStr, "test", "zs123")
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theID, err := rt.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is: %d\n", theID)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	} else {
		fmt.Println("初始化成功")
	}
	insertData()
}
