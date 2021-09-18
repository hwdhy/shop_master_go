package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/order"
	"shop_mater/utils"
	"time"
)

type RequestBody struct {
	AddressId  int    `json:"addressId"`
	PostScript string `json:"post_script"`
}

// Submit 提交订单
func Submit(c *gin.Context) {

	var rb RequestBody
	err := c.Bind(&rb)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var address models.Address
	db := database.DB.Model(models.Address{}).Where("id = ?", rb.AddressId).First(&address)
	if db.Error != nil {
		c.JSON(http.StatusOK, utils.ErrReturn(400, "请选择收货地址"))
		return
	}

	var carts []models.Cart
	tx := database.DB.Model(models.Cart{}).Debug().Where("user_id = ?", base.GetLoginUserID()).Where("session_id = ?", "1").Where("checked = 1").Find(&carts)
	if tx.Error != nil {
		c.JSON(http.StatusOK, utils.ErrReturn(400, "请选择商品"))
	}

	//var freightPrice = 0.0
	var goodsTotalPrice = 0.0

	for _, cart := range carts {
		goodsTotalPrice += float64(cart.Number) * cart.RetailPrice
	}

	var couponPrice float64
	orderTotalPrice := goodsTotalPrice
	//orderTotalPrice := goodsTotalPrice + freightPrice - couponPrice

	fmt.Println(orderTotalPrice)

	actualPrice := orderTotalPrice - 0
	currentTime := time.Now().Unix()

	orderInfo := models.Order{
		OrderSn:      utils.GenerateOrderNumber(),
		UserId:       base.GetLoginUserID(),
		Consignee:    address.Name,
		Mobile:       address.Mobile,
		Province:     address.ProvinceID,
		City:         address.CityID,
		District:     address.DistrictID,
		Address:      address.Address,
		FreightPrice: 0,
		Postscript:   rb.PostScript,
		CouponId:     0,
		CouponPrice:  couponPrice,
		AddTime:      currentTime,
		GoodsPrice:   goodsTotalPrice,
		OrderPrice:   orderTotalPrice,
		ActualPrice:  actualPrice,
	}
	create := database.DB.Model(models.Order{}).Create(&orderInfo)
	if create.Error != nil {
		c.JSON(http.StatusOK, utils.ErrReturn(400, "订单创建失败"))
		return
	}

	for _, item := range carts {
		orderGood := models.OrderGoods{
			OrderId:                     int(orderInfo.ID),
			GoodsId:                     item.GoodsId,
			GoodsSn:                     item.GoodsSn,
			ProductId:                   item.ProductID,
			GoodsName:                   item.GoodsName,
			ListPicUrl:                  item.ListPicUrl,
			MarketPrice:                 item.MarketPrice,
			RetailPrice:                 item.RetailPrice,
			Number:                      item.Number,
			GoodsSpecificationNameValue: item.GoodsSpecificationNameValue,
			GoodsSpecificationIds:       item.GoodsSpecificationIds,
		}
		database.DB.Debug().Model(models.OrderGoods{}).Create(&orderGood)
	}

	order.ClearBuyGoods(base.GetLoginUserID())

	c.JSON(http.StatusOK, utils.SuccessReturn(orderInfo))

}
