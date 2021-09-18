package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	order2 "shop_mater/service/order"
	"shop_mater/service/region"
	"shop_mater/utils"
	"strconv"
	"time"
)

type OrderInfo struct {
	models.Order
	ProvinceName       string `json:"province_name"`
	CityName           string `json:"city_name"`
	DistrictName       string `json:"district_name"`
	FullRegion         string `json:"full_region"`
	OrderStatusText    string `json:"order_status_text"`
	FormatAddTime      string `json:"add_time"`
	FormatFinalPayTime string `json:"final_pay_time"`
}

type OrderDetailRtnJson struct {
	OrderInfo    OrderInfo                `json:"orderInfo"`
	OrderGoods   []models.OrderGoods      `json:"orderGoods"`
	HandleOption order2.OrderHandleOption `json:"handleOption"`
}

// Detail 订单详情
func Detail(c *gin.Context) {

	orderId := c.DefaultQuery("orderId", "0")
	intOrderId, _ := strconv.Atoi(orderId)

	var order models.Order
	db := database.DB.Model(models.Order{}).Where("id = ?", intOrderId).Where("user_id = ?", base.GetLoginUserID()).First(&order)
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, "订单不存在")
		return
	}

	var orderInfo = OrderInfo{Order: order}

	orderInfo.ProvinceName = region.GetRegionName(order.Province)
	orderInfo.CityName = region.GetRegionName(order.City)
	orderInfo.DistrictName = region.GetRegionName(order.District)
	orderInfo.FullRegion = orderInfo.ProvinceName + orderInfo.CityName + orderInfo.DistrictName

	var orderGoods []models.OrderGoods
	database.DB.Model(models.OrderGoods{}).Where("order_id = ?", intOrderId).Find(&orderGoods)

	orderInfo.OrderStatusText = order2.GetOrderStatusText(intOrderId)
	orderInfo.FormatAddTime = time.Unix(orderInfo.AddTime, 0).Format("2006-01-02 15:04:05 PM")
	orderInfo.FormatFinalPayTime = time.Unix(time.Now().Add(time.Hour).Unix(), 0).Format("15:04")

	if orderInfo.OrderStatus == 0 {
		//todo 超时逻辑

	}

	handleOption := order2.GetOrderHandleOption(intOrderId)

	c.JSON(http.StatusOK, utils.SuccessReturn(OrderDetailRtnJson{
		OrderInfo:    orderInfo,
		OrderGoods:   orderGoods,
		HandleOption: handleOption,
	}))

}
