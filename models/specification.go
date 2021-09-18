package models

import (
	"gorm.io/gorm"
	"time"
)

// Specification 规格
type Specification struct {
	Name      string `json:"name"`       //名称
	SortOrder int    `json:"sort_order"` //排序

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Specification) TableName() string {
	return "specification"
}
