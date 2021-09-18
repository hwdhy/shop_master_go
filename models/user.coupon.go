package models

import (
	"gorm.io/gorm"
	"time"
)

// UserCoupon 用户优惠券
type UserCoupon struct {
	CouponID     int    `json:"coupon_id"`
	CouponNumber string `json:"coupon_number"`
	OrderID      int    `json:"order_id"`
	UsedTime     int    `json:"used_time"`
	UserID       int    `json:"user_id"`

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (UserCoupon) TableName() string {
	return "user_coupon"
}
