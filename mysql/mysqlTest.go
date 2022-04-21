package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "yoyo1033"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
	//userName = "fms_test"
	//password = "20g(981ogw*71t99)*(1kg"
	//ip       = "172.16.64.185"
	//port     = "3306"

	//dbName   = "fms_shop"
)

//Db数据库连接池
var DB *sql.DB

type Info struct {
	id    int64
	name  string
	age   int8
	sex   int8
	phone string
}

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func PingDB(db *sql.DB) {
	err := db.Ping()
	ErrorCheck(err)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//查询操作
func Query() {
	res, err := DB.Query("SELECT * FROM info where id >11")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var fo Info
		err := res.Scan(&fo.id, &fo.name, &fo.age, &fo.sex, &fo.phone)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", fo)
	}
	//返回json
	//jsons, err := json.Marshal(res)
	//if err != nil {
	//	fmt.Println("err",err)
	//}
	//fmt.Println(string(jsons))

}

func DeleteUser() bool {
	PingDB(DB)

	//delete

	stmt, e := DB.Prepare("delete from info where id =?")
	ErrorCheck(e)

	res, e := stmt.Exec("24")
	ErrorCheck(e)
	// affected rows
	a, e := res.RowsAffected()
	ErrorCheck(e)

	fmt.Println(a) // 1
	return true
}

func InsertUser() bool {
	sql := "insert into info(name,age,sex,phone) values('bao',10,1,1121)"

	res, err := DB.Exec(sql)

	if err != nil {
		panic(err.Error())
	}
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	fmt.Printf("this is id: %d\n", lastId)

	return true
}

func main() {
	InitDB()
	Query()
	//InsertUser()
	//DeleteUser()
	defer DB.Close()
}
