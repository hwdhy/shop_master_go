package models

import (
	"gorm.io/gorm"
	"time"
)

// Keywords 关键词
type Keywords struct {
	IsDefault int    `json:"is_default"`
	IsHot     int    `json:"is_hot"`
	IsShow    int    `json:"is_show"`
	Keyword   string `json:"keyword"`
	SchemeUrl string `json:"scheme_url"`
	SortOrder int    `json:"sort_order"`
	Type      int    `json:"type"`

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Keywords) TableName() string {
	return "keywords"
}
