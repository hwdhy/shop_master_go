package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/login"
	"shop_mater/service/weixin"
	"shop_mater/utils"
	"strconv"
	"time"
)

type AuthLoginBody struct {
	Code     string             `json:"code"`
	UserInfo weixin.ResUserInfo `json:"userInfo"`
}

// LoginByWeixin 微信登录
func LoginByWeixin(c *gin.Context) {
	//请求接口体
	var loginBody AuthLoginBody

	err := c.Bind(&loginBody)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	clientIP := c.ClientIP()

	//获取微信登录返回信息
	userInfo := weixin.Login(loginBody.Code, loginBody.UserInfo)

	//返回用户不为空   创建用户
	if userInfo != nil {
		var user models.User
		//判断用户是否已经存在，不存在则创建
		database.DB.Model(models.User{}).Where("weixin_openid = ?", userInfo.OpenID).First(&user)
		if user.ID == 0 {
			newUser := models.User{
				Username:      uuid.NewV4().String(),
				Password:      "",
				RegisterTime:  time.Now().Unix(),
				RegisterIp:    clientIP,
				Mobile:        "",
				WeixinOpenid:  userInfo.OpenID,
				Avatar:        userInfo.AvatarUrl,
				Gender:        userInfo.Gender,
				Nickname:      userInfo.NickName,
				LastLoginIp:   clientIP,
				LastLoginTime: time.Now().Unix(),
			}
			database.DB.Model(models.User{}).Create(&newUser)
			database.DB.Model(models.User{}).Where("weixin_openid = ?", userInfo.OpenID).First(&user)
		}

		userinfo := make(map[string]interface{})
		userinfo["id"] = user.ID
		userinfo["username"] = user.Username
		userinfo["nickname"] = user.Nickname
		userinfo["gender"] = user.Gender
		userinfo["avatar"] = user.Avatar
		userinfo["birthday"] = user.Birthday

		userIdString := strconv.Itoa(int(user.ID))
		sessionKey := login.Create(userIdString)

		rtnInfo := make(map[string]interface{})
		rtnInfo["token"] = sessionKey
		rtnInfo["userInfo"] = userinfo

		c.JSON(http.StatusOK, utils.SuccessReturn(rtnInfo))
	}
}
