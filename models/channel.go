package models

import (
	"gorm.io/gorm"
	"time"
)

// Channel 渠道
type Channel struct {
	Name      string `json:"name"`       //名称
	SortOrder int    `json:"sort_order"` //排序
	Url       string `json:"url"`        //路径                   *
	IconUrl   string `json:"icon_url"`   //图标路径                *

	ID        int64          `json:"id"`         //id             *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Channel) TableName() string {
	return "channel"
}
