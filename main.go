package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"shop_mater/database"
	"shop_mater/routers"
)

func main() {
	database.InitDatabase()

	r := gin.Default()
	//小程序端接口初始化
	routers.InitRouters(r)

	//后台接口初始化
	routers.AdminRouters(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
