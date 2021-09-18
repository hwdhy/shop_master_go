package models

import (
	"gorm.io/gorm"
	"time"
)

// OrderGoods 订单商品
type OrderGoods struct {
	GoodsId                     int     `json:"goods_id"`
	GoodsName                   string  `json:"goods_name"`
	GoodsSn                     string  `json:"goods_sn"`
	GoodsSpecificationIds       string  `json:"goods_specification_ids"`
	GoodsSpecificationNameValue string  `json:"goods_specification_name_value"`
	IsReal                      int     `json:"is_real"`
	ListPicUrl                  string  `json:"list_pic_url"`
	MarketPrice                 float64 `json:"market_price"`
	Number                      int     `json:"number"`
	OrderId                     int     `json:"order_id"`
	ProductId                   int     `json:"product_id"`
	RetailPrice                 float64 `json:"retail_price"`

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (OrderGoods) TableName() string {
	return "order_goods"
}
