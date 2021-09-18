package adminGoods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

// Create 创建产品
func Create(c *gin.Context) {

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
	goodsData.GoodsDesc = InputData.GoodsDesc
	goodsData.ListPicUrl = InputData.ListPicUrl
	goodsData.SortOrder = InputData.SortOrder

	if db := database.DB.Model(models.Goods{}).Create(&goodsData); db.Error != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, "")
}
