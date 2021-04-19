package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//初始化连接池
var db *sql.DB

type Db interface {
	Connect()
	Create()
	Query()

}

func Connect() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)//设置连接最大存活事件
	db.SetMaxOpenConns(10)//设置与数据库建立连接的最大数目
	db.SetMaxIdleConns(10)//设置连接池中的最大闲置连接数

	defer db.Close()

	err = db.Ping()//检测是否出现错误
	if err != nil {
		fmt.Printf("open %s invaild ,err'=", err)
		return
	}
	fmt.Println("Database connection successful")
	return
}

func Crete()  {

}
