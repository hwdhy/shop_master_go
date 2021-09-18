package models

import (
	"gorm.io/gorm"
	"time"
)

// SearchHistory 搜索记录
type SearchHistory struct {
	AddTime int64  `json:"add_time"`
	From    string `json:"from"`
	Keyword string `json:"keyword"`
	UserID  int    `json:"user_id"`

	ID        int            `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (SearchHistory) TableName() string {
	return "search_history"
}
