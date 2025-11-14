package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID       uint   `gorm:"primary_key;auto_increment"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func main() {
	//数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate(&User{})

	//服务器
	server := gin.Default()
	server.Use(cors.Default())
	//注册
	server.POST("/register", func(context *gin.Context) {
		u := User{}
		context.BindJSON(&u)
		//查询用户名是否存在
		res := db.Where("name = ?", u.Name).First(&User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册失败，用户名已存在!",
			})
		} else {
			db.Create(&u)
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册成功",
			})
		}
	})
	server.POST("/login", func(context *gin.Context) {
		u := User{}
		context.BindJSON(&u)
		res := db.Where("name = ? and password = ?", u.Name, u.Password).First(&User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "登录成功！",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "登录失败，用户名或密码错误！",
			})
		}
	})
	server.Run(":8080")

}
