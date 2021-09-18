package models

import (
	"gorm.io/gorm"
	"time"
)

// Ad 产品广告
type Ad struct {
	AdPositionId int       `json:"ad_position_id"`              //广告位置ID    *
	ImageUrl     string    `json:"image_url"`                   //图片路径      *
	Link         string    `json:"link"`                        //点击跳转链接   *
	MediaType    int8      `json:"media_type" gorm:"default:1"` //媒体类型
	Content      string    `json:"content"`                     //内容         *
	Enabled      int8      `json:"enabled"`                     //启用         *
	Name         string    `json:"name"`                        //名称
	EndTime      time.Time `json:"end_time"`                    //结束时间

	ID        int64          `json:"id"`         //id          *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Ad) TableName() string {
	return "ad"
}
