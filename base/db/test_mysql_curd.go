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

type user struct {
	id       int
	username string
	password string
}

func queryRow() {
	// 查单行
	sqlStr := "select id, username, password from user_tbl where id = ?"
	var u user
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d, name:%s, password:%s\n", u.id, u.username, u.password)
}

func queryMultiRow() {
	sqlStr := "select id, username, password from user_tbl where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query faild,err: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed,err: %v\n", err)
			return
		}
		fmt.Printf("id:%d, name:%s, password:%s\n", u.id, u.username, u.password)
	}

}

func updateData() {
	sql := "update user_tbl set username = ?, password = ? where id = ?"
	rt, err := db.Exec(sql, "kite2", "kite123", "2")
	if err != nil {
		fmt.Printf("更新失败,err: %v\n", err)
		return
	}

	rows, err := rt.RowsAffected()
	if err != nil {
		fmt.Printf("更新失败,err: %v\n", err)
		return
	}
	fmt.Printf("更新成功，更新的行数是: %d\n", rows)
}

func deleteData() {
	sql := "delete from user_tbl where id = ?"
	rt, err := db.Exec(sql, "1")
	if err != nil {
		fmt.Printf("删除,err: %v\n", err)
		return
	}

	rows, err := rt.RowsAffected()
	if err != nil {
		fmt.Printf("删除行失败,err: %v\n", err)
		return
	}
	fmt.Printf("删除成功，删除的行数是: %d\n", rows)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	} else {
		fmt.Println("初始化成功")
	}
	queryRow()
	queryMultiRow()
	updateData()
	deleteData()
}
