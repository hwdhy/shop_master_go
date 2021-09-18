package models

import (
	"gorm.io/gorm"
	"time"
)

// Gallery 商品视图
type Gallery struct {
	GoodsID   int    `json:"goods_id"`   //商品ID
	ImgUrl    string `json:"img_url"`    //图片URL                   *
	ImgDesc   string `json:"img_desc"`   //商品描述
	SortOrder int    `json:"sort_order"` //排序

	ID        int64          `json:"id"`         //id              *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Gallery) TableName() string {
	return "gallery"
}
