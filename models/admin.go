package models

import (
	"gorm.io/gorm"
	"time"
)

// Admin 管理后台账号
type Admin struct {
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Admin) TableName() string {
	return "admin"
}
