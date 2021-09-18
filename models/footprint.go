package models

import (
	"gorm.io/gorm"
	"time"
)

// Footprint 浏览记录
type Footprint struct {
	GoodsID int64 `json:"goods_id"` //商品ID                     *
	UserID  int64 `json:"user_id"`  //用户ID						*
	AddTime int64 `json:"add_time"` //添加时间					*

	ID        int64          `json:"id"`         //id			*
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Footprint) TableName() string {
	return "footprint"
}
