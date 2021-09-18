package models

import (
	"gorm.io/gorm"
	"time"
)

// GoodsSpecification 商品规格
type GoodsSpecification struct {
	GoodsId         int    `json:"goods_id"`         //商品ID
	PicUrl          string `json:"pic_url"`          //图片URL
	SpecificationID int    `json:"specification_id"` //规格ID
	Value           string `json:"value"`            //值

	ID        int            `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (GoodsSpecification) TableName() string {
	return "goods_specification"
}
