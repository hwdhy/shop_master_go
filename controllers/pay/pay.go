package pay

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/weixin"
	"shop_mater/utils"
	"strconv"
)

func Prepay(c *gin.Context) {
	orderId := c.DefaultQuery("orderId", "0")

	intOrderId, _ := strconv.Atoi(orderId)

	var order models.Order
	db := database.DB.Model(models.Order{}).Where("id = ?", intOrderId).First(&order)
	if db.Error != nil {
		c.JSON(400, "订单已取消")
		return
	}

	if order.PayStatus != 0 {
		c.JSON(400, "订单已支付，请不要重复操作!")
		return
	}

	var user models.User
	tx := database.DB.Model(models.User{}).Where("id = ?", order.UserId).First(&user)

	log.Println(user.WeixinOpenid)
	if tx.Error == nil && user.WeixinOpenid == "" {
		c.JSON(http.StatusOK, utils.ErrReturn(400, "微信支付失败！"))
		return
	}

	payInfo := weixin.PayInfo{
		OpenId:     user.WeixinOpenid,
		Body:       "order NO: " + order.OrderSn,
		OutTradeNo: order.OrderSn,
		TotalFee:   int64(order.ActualPrice * 100),
	}

	param, err := weixin.CreateUnifiedOrder(payInfo)
	fmt.Println(param)
	if err != nil {
		c.JSON(200, utils.ErrReturn(400, "微信支付失败"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(param))

}
