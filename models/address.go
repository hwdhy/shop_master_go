package models

import (
	"gorm.io/gorm"
	"time"
)

// Address 收货地址
type Address struct {
	Address    string `json:"address"`     //地址                *
	CityID     int    `json:"city_id"`     //城市ID
	CountryID  int    `json:"country_id"`  //国家ID
	DistrictID int    `json:"district_id"` //区域ID
	IsDefault  int    `json:"is_default"`  //是否默认地址           *
	Mobile     string `json:"mobile"`      //手机号码              *
	Name       string `json:"name"`        //名称                  *
	ProvinceID int    `json:"province_id"` //省份ID
	UserID     int    `json:"user_id"`     //用户ID

	ID        int64          `json:"id"`         //id             *
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Address) TableName() string {
	return "address"
}
