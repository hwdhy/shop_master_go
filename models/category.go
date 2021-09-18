package models

import (
	"gorm.io/gorm"
	"time"
)

// Category 分类
type Category struct {
	ImgUrl       string `json:"img_url"`        //图片URL
	IsShow       int    `json:"is_show"`        //是否显示 									*
	Keywords     string `json:"keywords"`       //关键词
	Level        string `json:"level"`          //等级
	Name         string `json:"name"`           //分类名称									*
	ParentId     int    `json:"parent_id"`      //父ID										*
	ShowIndex    int    `json:"show_index"`     //显示下标									*
	SortOrder    int    `json:"sort_order"`     //排序										*
	Type         int    `json:"type"`           //类型
	WapBannerUrl string `json:"wap_banner_url"` //网页横幅网址								*
	BannerUrl    string `json:"banner_url"`     //横幅网址
	FrontDesc    string `json:"front_desc"`     //前台描述									*
	FrontName    string `json:"front_name"`     //前台介绍                                   *
	IconUrl      string `json:"icon_url"`       //icon路径

	ID        int64          `json:"id"`         //id										*
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Category) TableName() string {
	return "category"
}
