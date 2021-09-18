package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

type CancelBody struct {
	OrderId string `json:"orderId"`
}

// Cancel 取消订单
func Cancel(c *gin.Context) {

	var cb CancelBody
	err := c.Bind(&cb)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	intOrderID, _ := strconv.Atoi(cb.OrderId)

	database.DB.Model(models.Order{}).Where("id = ?", intOrderID).Where("user_id = ?", base.GetLoginUserID()).Update("order_status", 101)

	c.JSON(http.StatusOK, utils.SuccessReturn("订单取消成功"))

}
