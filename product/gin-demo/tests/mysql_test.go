package tests

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysql(t *testing.T) {
	connstr := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		t.Fatal(err.Error())
	}

	//创建一张表
	db.Exec("create table person ")
	defer db.Close()
}
