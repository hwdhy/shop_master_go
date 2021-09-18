package models

import (
	"gorm.io/gorm"
	"time"
)

// Order 订单
type Order struct {
	Integral       int     `json:"integral"`        //不可缺少的
	IntegralMoney  float64 `json:"integral_money"`  //积分货币
	Mobile         string  `json:"mobile"`          //手机号
	OrderPrice     float64 `json:"order_price"`     //订单价格
	OrderSn        string  `json:"order_sn"`        //
	OrderStatus    int     `json:"order_status"`    //订单状态
	ParentId       int     `json:"parent_id"`       //父ID
	PayId          int     `json:"pay_id"`          //支付ID
	PayName        string  `json:"pay_name"`        //支付名称
	PayStatus      int     `json:"pay_status"`      //支付状态
	PayTime        int     `json:"pay_time"`        //支付时间
	Postscript     string  `json:"postscript"`      //后记
	ShippingFee    float64 `json:"shipping_fee"`    //运费
	ShippingStatus int     `json:"shipping_status"` //运输状态
	UserId         int     `json:"user_id"`         //用户ID
	ActualPrice    float64 `json:"actual_price"`    //实际价格
	AddTime        int64   `json:"add_time"`        //添加时间
	Address        string  `json:"address"`         //收货地址
	CallbackStatus string  `json:"callback_status"` //回调状态
	ConfirmTime    int     `json:"confirm_time"`    //确认时间
	Consignee      string  `json:"consignee"`       //收货人
	Country        int     `json:"country"`         //国家
	Province       int     `json:"province"`        //省
	City           int     `json:"city"`            //城市
	District       int     `json:"district"`        //区
	CouponId       int     `json:"coupon_id"`       //优惠券ID
	CouponPrice    float64 `json:"coupon_price"`    //优惠券价格
	FreightPrice   float64 `json:"freight_price"`   //运费价格
	GoodsPrice     float64 `json:"goods_price"`     //商品价格

	ID        int64          `json:"id"`         //id
	CreatedAt time.Time      `json:"created_at"` //创建时间
	UpdatedAt time.Time      `json:"updated_at"` //更新时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"` //删除时间
}

func (Order) TableName() string {
	return "order"
}
