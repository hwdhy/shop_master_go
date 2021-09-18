package login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/utils"
	"strings"
	"time"
)

var key = []byte("hwdhy-0426-0125")
var expireTime = 30
var LoginUserId string

var publicController = "index,catalog,topic,auth,goods,brand,search,region"
var publicAction = "comment/list,comment/count,cart/index,cart/add,cart/checked,cart/update,cart/delete,cart/goodscount,pay/notify"

// CustomClaims 自定义申明
type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Create 创建token
func Create(userID string) string {

	claims := CustomClaims{
		userID,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expireTime)).Unix(), //过期时间20分钟
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return ""
	}
	return tokenStr
}

// Parse 解析token
func Parse(tokenStr string) *jwt.Token {

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil
	}
	if token.Valid {
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("token is expired or not valid!")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return nil
}

// Middleware token校验  ----- 小程序端中间件
func Middleware(c *gin.Context) {
	controller, action := getControllerAndAction(c.Request.RequestURI)
	fmt.Println("c.Request.RequestURI:", action)
	fmt.Println("c.Request.RequestURI:", c.Request.Host)
	token := c.GetHeader("x-shop-token")
	if action == "auth/loginByWeixin" || action == "goods/count" || action == "index/index" {
		return
	}

	if token == "" {
		data := utils.ErrReturn(401, "need relogin")
		c.Abort()
		c.JSON(http.StatusOK, data)
		c.Redirect(200, "/")
		return
	}
	//获取用户ID
	LoginUserId = GetUserID(token)
	//todo 获取配置文件，暂时存string数组中
	if !strings.Contains(publicController, controller) && !strings.Contains(publicAction, action) {
		if LoginUserId == "" {
			data := utils.ErrReturn(401, "need relogin")
			c.Abort()
			c.JSON(http.StatusOK, data)
			c.Redirect(http.StatusOK, "/")
		}
	}
}

// AdminMiddleware 后端中间件
func AdminMiddleware(c *gin.Context) {
	_, action := getControllerAndAction(c.Request.RequestURI)
	fmt.Println("c.Request.RequestURI:", action)
	token := c.GetHeader("token")
	fmt.Println("token============", token)
	if action == "user/login" {
		return
	}

	if token == "" {
		c.Abort()
		c.JSON(http.StatusUnauthorized, "请先登录")
	}

	//获取用户名称
	username := GetUserID(token)

	fmt.Println("username==========", username)

	if username == "" {
		c.Abort()
		c.JSON(http.StatusUnauthorized, "验证失败")
	}
}

//获取请求action
func getControllerAndAction(rawValue string) (controller, action string) {
	vals := strings.Split(rawValue, "/")

	return vals[2], vals[2] + "/" + vals[3]
}

// GetUserID 获取用户ID
func GetUserID(tokenStr string) string {
	token := Parse(tokenStr)
	if token == nil {
		return ""
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims.UserID
	}
	return ""
}
