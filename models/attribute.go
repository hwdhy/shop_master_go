package models

import (
	"gorm.io/gorm"
	"time"
)

// Attribute 属性
type Attribute struct {
	InputType           int    `json:"input_type"`            //输入类型
	Name                string `json:"name"`                  //名称
	SortOrder           int    `json:"sort_order"`            //排序
	Values              string `json:"values"`                //值
	AttributeCategoryID int    `json:"attribute_category_id"` //属性分类ID

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Attribute) TableName() string {
	return "attribute"
}
