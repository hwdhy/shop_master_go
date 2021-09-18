package models

import (
	"gorm.io/gorm"
	"time"
)

// GoodsAttribute 商品属性--属性 关联表
type GoodsAttribute struct {
	AttributeID int    `json:"attribute_id"` //属性ID
	GoodsID     int    `json:"goods_id"`     //商品ID
	Value       string `json:"value"`        //值

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (GoodsAttribute) TableName() string {
	return "goods_attribute"
}
