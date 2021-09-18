package models

import (
	"gorm.io/gorm"
	"time"
)

type Topic struct {
	IsShow          int8   `json:"is_show"`           //是否显示
	ItemPicUrl      string `json:"item_pic_url"`      //图片列表
	PriceInfo       string `json:"price_info"`        //图片信息
	ReadCount       string `json:"read_count"`        //阅读数量
	ScenePicUrl     string `json:"scene_pic_url"`     //场景图片网址
	SortOrder       int64  `json:"sort_order"`        //排序
	Subtitle        string `json:"subtitle"`          //副标题
	Title           string `json:"title"`             //标题
	TopicCategoryId int64  `json:"topic_category_id"` //专题分类ID
	TopicTagId      int64  `json:"topic_tag_id"`      //专题标签ID
	TopicTemplateId int64  `json:"topic_template_id"` //专题模板ID
	Avatar          string `json:"avatar"`            //头像
	Content         string `json:"content"`           //内容

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Topic) TableName() string {
	return "topic"
}
