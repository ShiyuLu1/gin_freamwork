package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"github.com/go-redis/redis/v8"
	"context"
)

//数据库的基础信息
const (
	userName = 
	password = ""
	ip = ""
	port = ""
	dbName = ""
)

func InitMysql() *sql.DB {

	//mysql 数据库
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ := sql.Open("mysql", path)
	if DB == nil {
		log.Fatal("连接失败！")
		return nil
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(5)
	//验证连接
	if err := DB.Ping(); err != nil{
		log.Fatal("opon database fail")
		return nil
	}
	return DB
}

func InitRedis() *redis.Client  {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  "",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	result := rdb.Ping(context.Background())
	fmt.Println("redis ping:", result.Val())
	if result.Val()!="PONG"{
		// 连接有问题
		return nil
	}
	return rdb
}
