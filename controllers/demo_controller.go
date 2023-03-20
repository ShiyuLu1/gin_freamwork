package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_mvc/models"
	"go_mvc/conf/database"
	"log"
	"database/sql"
	"github.com/go-redis/redis/v8"
	"fmt"
)

var (
	db *sql.DB
	rd *redis.Client
)

func init() {
	log.Println(">>>> get database connection start <<<<")
	db = database.InitMysql()
	rd = database.InitRedis()
}

type Result struct {
	Code  int 
	Message string 
	Data  []models.User 
}
func GetUsers(c *gin.Context) {
	key := "111"
	val := "222"
	ret, err := rd.Set(c, key,val,0).Result()
	fmt.Println("err:", err)
	fmt.Println(ret)
	//c.JSON(200, gin.H{"result": result})
	rows, err := db.Query("select id, username, password from user_info")
	if err != nil {
		log.Fatal(err.Error())
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	result := &Result{
		Code:  0,
		Message: "success",
		Data: users,
	}
	c.JSON(http.StatusOK, result)
}

