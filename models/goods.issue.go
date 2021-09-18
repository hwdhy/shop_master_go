package models

import (
	"gorm.io/gorm"
	"time"
)

// GoodsIssue 商品问题
type GoodsIssue struct {
	Answer   string `json:"answer"`   //回答						*
	GoodsID  string `json:"goods_id"` //商品ID					*
	Question string `json:"question"` //问题						*

	ID        int64          `json:"id"`         //id            *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}
