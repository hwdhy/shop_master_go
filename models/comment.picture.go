package models

import (
	"gorm.io/gorm"
	"time"
)

// CommentPicture 评论图片
type CommentPicture struct {
	CommentID int64  `json:"comment_id"` //评论ID
	PicUrl    string `json:"pic_url"`    //图片url                   *
	SortOrder int64  `json:"sort_order"` //排序

	ID        int64          `json:"id"`         //id               *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (CommentPicture) TableName() string {
	return "comment_picture"
}
