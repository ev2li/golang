package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func HandlerError(err error, why string) {
	if err != nil {
		fmt.Println("ERROR OCCURED!!! err = ", err, why)
	}
}

//将文本大数据入库，成功后做一个文件标记
func init() {
	//判断是否已经初始化数据库文件了
	_, err := os.Stat("/tmp/db_ok.mark")
	if err == nil {
		return
	}
	//打开数据库
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/kaifang2")
	HandlerError(err, "sqlx.Open")
	defer db.Close()

	//必要时建表
	_, err = db.Exec("create table if not exists kfperson(id int primary key auto_increment,name varchar(20),idcard char(18));")
	HandlerError(err, "db.Exec(crate table)")

	//分批次读入文本文件
	file, err := os.Open("/tmp/haha.txt")
	HandlerError(err, "os.Open")
	reader := bufio.NewReader(file)

	for {
		lineBytes, _, err := reader.ReadLine()
		HandlerError(err, "reader.ReadLine")
		if err == io.EOF {
			break
		}

		//逐条入库
		lineStr := string(lineBytes)
		fields := strings.Split(lineStr, ",")
		name, idcard := fields[0], fields[1]
		rt, err := db.Exec("Insert into kfperson(name, idcard) values(?, ?);", name, idcard)
		HandlerError(err, "db.Exec")
		if n, err := rt.RowsAffected(); err == nil && n > 0 {
			fmt.Printf("插入%s成功\n", name)
			if lastId, err := rt.LastInsertId(); err == nil && lastId > 10000 {
				break
			}
		}
	}
	fmt.Println("数据初始化成功")

	//创建一个标记文件
	_, err = os.Create("tmp/db_ok.mark")
	if err == nil {
		fmt.Println("数据库标记文件已创建")
	}
}

var (
	kfMap map[string][]kfperson
)

func main() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/kaifang2")
	HandlerError(err, "sqlx.Open")
	defer db.Close()
	//初始化缓存
	kfMap = make(map[string][]KfPerson, 0)

	for {
		var name string
		//接收用户想要查询的姓名
		fmt.Print("请输入要查询的开房者的姓名:")
		fmt.Scanf("%s", &name)

		if name == "exit" {
			break
		}
		//先查缓存
		if value, ok := kfMap[name]; ok {
			fmt.Println(value)
			continue
		}

		//查DB，结果放入缓冲
		kfPeople := make([]KfPerson, 0)
		err = db.Select(&kfPeople, "select name, idcard from kfperson where name like ?;", name)
		HandlerError(err, "db.Select")
		fmt.Printf("kfPeopel: %v\n", kfPeople)

		//入内存

		kfMap[name] = kfPeople
	}
}
