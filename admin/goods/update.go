package adminGoods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

// Update 产品修改
func Update(c *gin.Context) {

	var InputData models.Goods
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var goodsData models.Goods
	goodsData.Name = InputData.Name
	goodsData.RetailPrice = InputData.RetailPrice
	goodsData.GoodsNumber = InputData.GoodsNumber
	goodsData.IsHot = InputData.IsHot
	goodsData.IsNew = InputData.IsNew
	goodsData.GoodsBrief = InputData.GoodsBrief
	goodsData.ListPicUrl = InputData.ListPicUrl
	goodsData.GoodsDesc = InputData.GoodsDesc
	goodsData.SortOrder = InputData.SortOrder

	if db := database.DB.Model(models.Goods{}).Where("id = ?", InputData.ID).Updates(&InputData); db.Error != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, "")
}
