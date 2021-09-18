package models

import (
	"gorm.io/gorm"
	"time"
)

// Region 地区
type Region struct {
	AgencyId int    `json:"agency_id"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
	Type     int    `json:"type"`

	ID        int            `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Region) TableName() string {
	return "region"
}
