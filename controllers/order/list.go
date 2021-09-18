package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/order"
	"shop_mater/utils"
)

// OrderListRtnJson 订单列表返回数据
type OrderListRtnJson struct {
	models.Order
	GoodsList       []models.OrderGoods     `json:"goods_list"`
	GoodsCount      int                     `json:"goods_count"`
	OrderStatusText string                  `json:"order_status_text"`
	HandOption      order.OrderHandleOption `json:"hand_option"`
}

// List 订单列表
func List(c *gin.Context) {
	var orders []models.Order
	database.DB.Model(models.Order{}).Where("user_id = ?", base.GetLoginUserID()).Find(&orders)

	firstPageDataOrders := GetOrderPageData(orders, 1, 10)

	var rtnOrderList []OrderListRtnJson
	var orderGoods []models.OrderGoods

	db := database.DB.Model(models.OrderGoods{})

	for _, item := range firstPageDataOrders.Data.([]models.Order) {
		db.Where("order_id = ?", item.ID).Find(&orderGoods)

		var goodsCount int
		for _, item := range orderGoods {
			goodsCount += item.Number
		}

		orderStatusText := order.GetOrderStatusText(int(item.ID))
		orderHandleOption := order.GetOrderHandleOption(int(item.ID))

		var orderListRtn = OrderListRtnJson{
			Order:           item,
			GoodsList:       orderGoods,
			GoodsCount:      goodsCount,
			OrderStatusText: orderStatusText,
			HandOption:      orderHandleOption,
		}
		rtnOrderList = append(rtnOrderList, orderListRtn)
	}
	firstPageDataOrders.Data = rtnOrderList

	c.JSON(http.StatusOK, utils.SuccessReturn(firstPageDataOrders))

}
