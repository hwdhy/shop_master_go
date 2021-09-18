package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type CountReturnData struct {
	GoodsCount int64 `json:"goods_count"`
}

// Count 商品数量
func Count(c *gin.Context) {

	var Total int64
	database.DB.Model(models.Goods{}).Where("is_delete = ?", 0).Where("is_on_sale = ?", 1).Count(&Total)

	countRtn := CountReturnData{
		GoodsCount: Total,
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(countRtn))
}
