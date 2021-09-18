package models

import (
	"gorm.io/gorm"
	"time"
)

// Product 产品
type Product struct {
	GoodsId               int     `json:"goods_id"`                //商品ID
	GoodsNumber           int     `json:"goods_number"`            //商品货号
	GoodsSn               string  `json:"goods_sn"`                //
	GoodsSpecificationIds string  `json:"goods_specification_ids"` //商品规格
	RetailPrice           float64 `json:"retail_price"`            //零售价

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Product) TableName() string {
	return "product"
}
