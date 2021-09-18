package models

import (
	"gorm.io/gorm"
	"time"
)

// Comment 评论
type Comment struct {
	MewContent string `json:"mew_content"` //内容                         *
	Status     int8   `json:"status"`      //状态
	TypeID     int8   `json:"type_id"`     //类型
	UserID     int    `json:"user_id"`     //用户ID
	ValueID    int    `json:"value_id"`    //值ID

	AddTime int64  `json:"add_time"` //添加时间                           *
	Content string `json:"content"`  //内容                              *

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Comment) TableName() string {
	return "comment"
}
