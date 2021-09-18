package models

import (
	"gorm.io/gorm"
	"time"
)

// AdPosition 广告位置
type AdPosition struct {
	Name   string `json:"name"`   //名称
	Desc   string `json:"desc"`   //描述
	Height int    `json:"height"` //高度
	Width  int    `json:"width"`  //宽度

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (AdPosition) TableName() string {
	return "ad_position"
}
