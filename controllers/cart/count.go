package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/utils"
)

// GoodsCount 商品件数
type GoodsCount struct {
	CartTotal CartTotal `json:"cartTotal"`
}

// Count 购物车商品件数
func Count(c *gin.Context) {
	cartData := GetCart()

	goodsCount := GoodsCount{
		CartTotal: CartTotal{
			GoodsCount: cartData.CartTotal.GoodsCount,
		},
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(goodsCount))
}
