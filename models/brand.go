package models

import (
	"gorm.io/gorm"
	"time"
)

// Brand 品牌
type Brand struct {
	IsNew         int8    `json:"is_new"`           //是否最新
	IsShow        int8    `json:"is_show"`          //是否显示
	ListPicUrl    string  `json:"list_pic_url"`     //图片URL列表
	Name          string  `json:"name"`             //名称                *
	NewPicUrl     string  `json:"new_pic_url"`      //最新图片
	NewSortOrder  int8    `json:"new_sort_order"`   //最新排序
	PicUrl        string  `json:"pic_url"`          //图片URL
	SimpleDesc    string  `json:"simple_desc"`      //简介
	SortOrder     int8    `json:"sort_order"`       //排序
	AppListPicUrl string  `json:"app_list_pic_url"` //App图片列表          *
	FloorPrice    float64 `json:"floor_price"`      //底价                *

	ID        int64          `json:"id"`         //id                    *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Brand) TableName() string {
	return "brand"
}
