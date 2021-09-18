package order

import (
	"shop_mater/database"
	"shop_mater/models"
)

type OrderHandleOption struct {
	Cancel   bool `json:"cancel"`
	Delete   bool `json:"delete"`
	Pay      bool `json:"pay"`
	Comment  bool `json:"comment"`
	Delivery bool `json:"delivery"`
	Confirm  bool `json:"confirm"`
	Return   bool `json:"return"`
	Buy      bool `json:"buy"`
}

func GetOrderStatusText(orderId int) string {

	var order models.Order
	database.DB.Model(models.Order{}).Where("id = ?", orderId).First(&order)
	// 订单流程：下单成功－》支付订单－》发货－》收货－》评论
	// 订单相关状态字段设计，采用单个字段表示全部的订单状态
	// 1xx表示订单取消和删除等状态 0订单创建成功等待付款，101订单已取消，102订单已删除
	// 2xx表示订单支付状态,201订单已付款，等待发货
	// 3xx表示订单物流相关状态,300订单已发货，301用户确认收货
	// 4xx表示订单退换货相关的状态,401没有发货，退款402,已收货，退款退货
	// 如果订单已经取消或是已完成，则可删除和再次购买
	var statusText = "未付款"
	switch order.OrderStatus {
	case 0:
		statusText = "未付款"
	case 101:
		statusText = "订单已取消"
	case 102:
		statusText = "订单已删除"
	case 201:
		statusText = "订单已付款"
	case 301:
		statusText = "订单已发货"
	case 302:
		statusText = "用户确认收货"

	}
	return statusText
}

func GetOrderHandleOption(orderId int) OrderHandleOption {
	// 订单流程：下单成功－》支付订单－》发货－》收货－》评论
	// 订单相关状态字段设计，采用单个字段表示全部的订单状态
	// 1xx表示订单取消和删除等状态 0订单创建成功等待付款，101订单已取消，102订单已删除
	// 2xx表示订单支付状态,201订单已付款，等待发货
	// 3xx表示订单物流相关状态,300订单已发货，301用户确认收货
	// 4xx表示订单退换货相关的状态,401没有发货，退款402,已收货，退款退货
	// 如果订单已经取消或是已完成，则可删除和再次购买

	var handoption = OrderHandleOption{
		Cancel:   false,
		Delete:   false,
		Pay:      false,
		Comment:  false,
		Delivery: false,
		Confirm:  false,
		Return:   false,
		Buy:      false,
	}

	var order models.Order
	database.DB.Model(models.Order{}).Where("id = ?", orderId).First(&order)

	switch order.OrderStatus {
	case 0:
		handoption.Cancel = true
		handoption.Pay = true
	case 101:
		handoption.Delete = true
		handoption.Buy = true
	case 201:
		handoption.Return = true
	case 300:
		handoption.Cancel = true
		handoption.Pay = true
		handoption.Return = true
	case 301:
		handoption.Delete = true
		handoption.Comment = true
		handoption.Buy = true
	}
	return handoption
}
