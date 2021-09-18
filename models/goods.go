package models

import (
	"gorm.io/gorm"
	"time"
)

// Goods 商品
type Goods struct {
	IsAppExclusive    int8    `json:"is_app_exclusive"`    //是否APP独有
	IsDelete          int8    `json:"is_delete"`           //是否删除
	IsHot             int8    `json:"is_hot"`              //是否最热           *
	IsLimited         int64   `json:"is_limited"`          //是否有限
	IsNew             int8    `json:"is_new"`              //是否最新			   *
	IsOnSale          int8    `json:"is_on_sale"`          //是否在售
	Keywords          string  `json:"keywords"`            //关键词
	ListPicUrl        string  `json:"list_pic_url"`        //图片列表            *
	Name              string  `json:"name"`                //产品名称				*
	PrimaryPicUrl     string  `json:"primary_pic_url"`     //主图路径             *
	PrimaryProductId  int64   `json:"primary_product_id"`  //主要产品ID
	PromotionDesc     string  `json:"promotion_desc"`      //促销描述
	PromotionTag      string  `json:"promotion_tag"`       //促销标签
	RetailPrice       float64 `json:"retail_price"`        //零售价				*
	SellVolume        int64   `json:"sell_volume"`         //销售量
	SortOrder         int64   `json:"sort_order"`          //排序
	UnitPrice         float64 `json:"unit_price"`          //单价
	AppExclusivePrice string  `json:"app_exclusive_price"` //app独家价格
	AttributeCategory int64   `json:"attribute_category"`  //类别属性
	BrandID           int64   `json:"brand_id"`            //品牌ID
	CategoryId        int64   `json:"category_id"`         //分类ID
	CounterPrice      string  `json:"counter_price"`       //还价
	ExtraPrice        string  `json:"extra_price"`         //额外价格
	GoodsBrief        string  `json:"goods_brief"`         //商品简介				*
	GoodsDesc         string  `json:"goods_desc"`          //商品描述				*
	GoodsNumber       int64   `json:"goods_number"`        //商品数量				*
	GoodsSn           string  `json:"goods_sn"`
	GoodsUnit         string  `json:"goods_unit"` //商品单位

	ID        int64          `json:"id"`         //id							*
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Goods) TableName() string {
	return "goods"
}
