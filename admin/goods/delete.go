package adminGoods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type GoodsDeleteInput struct {
	ID int64 `json:"id"` //商品ID
}

// Delete 后台商品删除接口
func Delete(c *gin.Context) {

	var InputData GoodsDeleteInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	database.DB.Where("id = ?", InputData.ID).Delete(&models.Goods{})

	c.JSON(http.StatusOK, "")
}
