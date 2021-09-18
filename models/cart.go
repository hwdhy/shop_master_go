package models

import (
	"gorm.io/gorm"
	"time"
)

// Cart 购物车
type Cart struct {
	ListPicUrl                  string  `json:"list_pic_url"`                   //图片路径          *
	MarketPrice                 float64 `json:"market_price"`                   //市场价
	Number                      int     `json:"number"`                         //商品数量          *
	ProductID                   int     `json:"product_id"`                     //产品ID
	RetailPrice                 float64 `json:"retail_price"`                   //零售价            *
	SessionID                   string  `json:"session_id"`                     //sessionID
	UserID                      int     `json:"user_id"`                        //用户ID
	Checked                     int     `json:"checked"`                        //选中               *
	GoodsId                     int     `json:"goods_id"`                       //商品ID
	GoodsName                   string  `json:"goods_name"`                     //商品名称            *
	GoodsSn                     string  `json:"goods_sn"`                       //
	GoodsSpecificationIds       string  `json:"goods_specification_ids"`        //商品规格ID
	GoodsSpecificationNameValue string  `json:"goods_specification_name_value"` //商品规格名称         *

	ID        int64          `json:"id"`         //id                                                *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间

}

func (Cart) TableName() string {
	return "cart"
}
