package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type IndexCartData struct {
	CartList  []models.Cart `json:"cart_list"`
	CartTotal CartTotal     `json:"cart_total"`
}

type CartTotal struct {
	GoodsCount         int64 `json:"goods_count"`
	GoodsAmount        float64 `json:"goods_amount"`
	CheckedGoodsCount  int64 `json:"checked_goods_count"`
	CheckedGoodsAmount float64 `json:"checked_goods_amount"`
}

// GetCart 购物车数据
func GetCart() IndexCartData {

	var cartData []models.Cart

	database.DB.Model(models.Cart{}).Debug().
		Where("user_id", base.GetLoginUserID()).
		Where("session_id = ?", "1").
		Find(&cartData)

	var goodsCount int64
	var goodsAmount float64
	var checkedGoodsCount int64
	var checkedGoodsAmount float64

	for _, item := range cartData {
		goodsCount += int64(item.Number)
		goodsAmount += float64(item.Number) * item.RetailPrice
		if item.Checked == 1 {
			checkedGoodsCount += int64(item.Number)
			checkedGoodsAmount += float64(item.Number) * item.RetailPrice
		}

		var goodsData models.Goods
		database.DB.Model(models.Goods{}).Where("id = ?", item.GoodsId).First(&goodsData)

		item.ListPicUrl = goodsData.ListPicUrl
	}

	return IndexCartData{
		CartList: cartData,
		CartTotal: CartTotal{
			GoodsCount:         goodsCount,
			GoodsAmount:        goodsAmount,
			CheckedGoodsCount:  checkedGoodsCount,
			CheckedGoodsAmount: checkedGoodsAmount,
		},
	}
}

// Index 购物车数据
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))
}
