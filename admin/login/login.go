package adminLogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/login"
)

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 后台登录
func Login(c *gin.Context) {

	fmt.Println("请求头：", c.Request.Header)

	var lr LoginRequest
	err := c.Bind(&lr)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	var adminData models.Admin
	//判断用户是否存在
	database.DB.Model(models.Admin{}).Where("username = ? and password = ?", lr.Username, lr.Password).First(&adminData)

	if adminData.ID == 0 {
		c.JSON(http.StatusBadRequest, "用户不存在")
		return
	}
	//生成token返回
	token := login.Create(lr.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
