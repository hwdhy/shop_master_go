package models

import (
	"gorm.io/gorm"
	"time"
)

// Collect 收藏
type Collect struct {
	IsAttention int   `json:"is_attention"` //是否点赞
	TypeId      int   `json:"type_id"`      //类型ID
	UserId      int   `json:"user_id"`      //用户ID
	ValueId     int   `json:"value_id"`     //值ID
	AddTime     int64 `json:"add_time"`     //新增时间

	ID        int64          `json:"id"`         //id              *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Collect) TableName() string {
	return "collect"
}
