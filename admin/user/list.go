package adminUser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"time"
)

type UserListInput struct {
	Current  int `json:"current"`
	PageSize int `json:"page_size"`
}

type UserListOutput struct {
	Current  int                  `json:"current"`
	PageSize int                  `json:"page_size"`
	Total    int64                `json:"total"`
	Data     []UserListOutputData `json:"data"`
}

type UserListOutputData struct {
	ID            int    `json:"id"`
	NickName      string `json:"nick_name"`
	WeixinOpenid  string `json:"weixin_openid"`
	Avatar        string `json:"avatar"`
	RegisterTime  string `json:"register_time"`
	Mobile        string `json:"mobile"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime string `json:"last_login_time"`
}

// List 用户列表
func List(c *gin.Context) {
	var InputData UserListInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("bind:", err)
		return
	}

	Offset := (InputData.Current - 1) * InputData.PageSize
	var Total int64
	var userData []models.User

	database.DB.Debug().Model(models.User{}).Count(&Total).Offset(Offset).Limit(InputData.PageSize).Find(&userData)

	var resData UserListOutputData
	var resDatas []UserListOutputData
	for _, item := range userData {
		resData.ID = int(item.ID)
		resData.NickName = item.Nickname
		resData.WeixinOpenid = item.WeixinOpenid
		resData.Avatar = item.Avatar
		resData.RegisterTime = time.Unix(item.RegisterTime, 0).Format("2006-01-02 15:04:05")
		resData.Mobile = item.Mobile
		resData.LastLoginIp = item.LastLoginIp
		resData.LastLoginTime = time.Unix(item.LastLoginTime, 0).Format("2006-01-02 15:04:05")

		resDatas = append(resDatas, resData)
	}

	c.JSON(http.StatusOK, UserListOutput{
		Current:  InputData.Current,
		PageSize: InputData.PageSize,
		Total:    Total,
		Data:     resDatas,
	})
}
