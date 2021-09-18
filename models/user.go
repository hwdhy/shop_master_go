package models

import (
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	LastLoginIp   string `json:"last_login_ip"`   //最后登录IP       *
	LastLoginTime int64  `json:"last_login_time"` //最后登录时间      *
	Mobile        string `json:"mobile"`          //手机号			*
	Nickname      string `json:"nickname"`        //昵称
	Password      string `json:"password"`        //密码
	RegisterIp    string `json:"register_ip"`     //注册IP
	RegisterTime  int64  `json:"register_time"`   //注册时间         *
	UserLevelId   int    `json:"user_level_id"`   //用户等级
	Username      string `json:"username"`        //用户名			*
	WeixinOpenid  string `json:"weixin_openid"`   //微信ID			*
	Avatar        string `json:"avatar"`          //头像				*
	Birthday      int    `json:"birthday"`        //生日
	Gender        int    `json:"gender"`          //性别

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (User) TableName() string {
	return "user"
}
