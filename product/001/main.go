package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const CACHE_LEN = 2

var (
	kfMap map[string]QueryResult
	// chanSema chan int
	chanData chan *KfPerson
	db       *sqlx.DB
)

func HandlerError(err error, why string) {
	if err != nil {
		fmt.Println("ERROR OCCURED!!! err = ", err, why)
	}
}

//将文本大数据入库，成功后做一个文件标记
func init() {
	var err error
	//判断是否已经初始化数据库文件了
	_, err = os.Stat("/tmp/db_ok.mark")
	if err == nil {
		return
	}
	//打开数据库

	db, err = sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/kaifang2")
	HandlerError(err, "sqlx.Open")
	defer db.Close()

	//必要时建表
	_, err = db.Exec("create table if not exists kfperson(id int primary key auto_increment,name varchar(20),idcard char(18));")
	HandlerError(err, "db.Exec(crate table)")

	//信号里管道(控制并发数)
	// chanSema = make(chan int, 100)
	chanData = make(chan *KfPerson, 10000000)
	//开辟协程，从数据管道获取信息，插入数据库
	for i := 0; i < 100; i++ {
		go insertKfPerson()
	}

	file, err := os.Open("/tmp/haha.txt")
	HandlerError(err, "os.Open")
	reader := bufio.NewReader(file)
	//分批次读入文本文件
	for {
		lineBytes, _, err := reader.ReadLine()
		HandlerError(err, "reader.ReadLine")
		if err == io.EOF {
			//关闭数据管道
			close(chanData)
			break
		}

		//逐条入库
		lineStr := string(lineBytes)
		fields := strings.Split(lineStr, ",")
		name, idcard := fields[0], fields[1]
		//抛弃过长的名字
		name = strings.TrimSpace(name)
		if len(strings.Split(name, "")) > 20 {
			continue
		}

		kfPerson := KfPerson{Name: name, Idcard: idcard}
		//方案一:开2000w协程行不通，耗尽资源，程序崩溃
		// go insertKfPerson(db, &kfPerson)
		//方案二：开有限协程，从管道中读取数据
		chanData <- &kfPerson

	}
	fmt.Println("数据初始化成功")

	//创建一个标记文件
	_, err = os.Create("tmp/db_ok.mark")
	if err == nil {
		fmt.Println("数据库标记文件已创建")
	}

}

func insertKfPerson() {
	for kfPerson := range chanData {
		rt, err := db.Exec("Insert into kfperson(name, idcard) values(?, ?);", kfPerson.Name, kfPerson.Idcard)
		HandlerError(err, "db.Exec")
		if n, err := rt.RowsAffected(); err == nil && n > 0 {
			fmt.Printf("插入%s成功\n", kfPerson.Name)
		}
	}

}

func main() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/kaifang2")
	HandlerError(err, "sqlx.Open")
	defer db.Close()
	//初始化缓存
	kfMap = make(map[string]QueryResult, 0)

	for {
		var name string
		//接收用户想要查询的姓名
		fmt.Print("请输入要查询的开房者的姓名:")
		fmt.Scanf("%s", &name)

		if name == "exit" {
			break
		}

		//查看所有缓存的key
		if name == "cache" {
			fmt.Printf("共缓存了%d条结果\n", len(kfMap))
			for key := range kfMap {
				fmt.Println(key)
			}
		}
		//先查缓存
		if qr, ok := kfMap[name]; ok {
			fmt.Println("查询到%d条结果:", len(qr.value))
			fmt.Println(qr.value)
			continue
		}

		//查DB，结果放入缓冲
		kfPeople := make([]KfPerson, 0)
		err = db.Select(&kfPeople, "select name, idcard from kfperson where name like ?;", name)
		HandlerError(err, "db.Select")
		fmt.Println("查询到%d条结果:", len(kfPeople))
		fmt.Printf("kfPeopel: %v\n", kfPeople)

		//入内存
		queryResult := QueryResult{value: kfPeople}
		queryResult.cacheTime = time.Now().UnixNano()
		queryResult.count = 1
		kfMap[name] = queryResult
		//有必要时淘汰一些缓存

		if len(kfMap) > CACHE_LEN {
			delkey := UpdateCache(&kfMap)
			fmt.Printf("delkey: %v已被淘汰出缓存\n", delkey)
		}
	}
}

//整理缓存,删除加入最早的
func UpdateCache(cacheMap *map[string]QueryResult) (delkey string) {
	earliestTime := time.Now().UnixNano()
	for key, value := range *cacheMap {
		if value.cacheTime < earliestTime {
			earliestTime = value.cacheTime
			delkey = key
		}
	}
	delete(*cacheMap, delkey)
	return delkey
}
