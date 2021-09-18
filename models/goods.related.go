package models

import (
	"gorm.io/gorm"
	"time"
)

// GoodsRelated 关联商品
type GoodsRelated struct {
	GoodsID        int64 `json:"goods_id"`         //商品ID         *
	RelatedGoodsID int64 `json:"related_goods_id"` //关联商品ID      *

	ID        int            `json:"id"`         //id				*
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间

}

func (GoodsRelated) TableName() string {
	return "goods_related"
}
