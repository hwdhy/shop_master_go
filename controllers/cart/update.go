package cart

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strings"
)

type CartUpdateBody struct {
	GoodsId   int `json:"goodsId"`
	ProductId int `json:"productId"`
	Number    int `json:"number"`
	Id        int `json:"id"`
}

type GoodsSpecification struct {
	models.GoodsSpecification
	Name string
}

// Update 更新购物车
func Update(c *gin.Context) {

	var updateBody CartUpdateBody
	err := c.Bind(&updateBody)
	if err != nil {
		println("err:", err)
		return
	}

	intGoodsID := updateBody.GoodsId
	intProductID := updateBody.ProductId
	intNumber := updateBody.Number
	intID := updateBody.Id

	var productData models.Product
	db := database.DB.Model(models.Product{}).
		Where("goods_id = ?", intGoodsID).Where("id = ?", intProductID).First(&productData)

	if db.Error != nil || productData.GoodsNumber < intNumber {
		c.JSON(http.StatusBadRequest, "库存不足")
		return
	}

	var CartData models.Cart
	database.DB.Model(models.Cart{}).Where("id = ?", intID).First(&CartData)

	if CartData.ProductID == intProductID {
		database.DB.Model(models.Cart{}).Where("id = ?", CartData.ID).Update("number", intNumber)

		c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))
		return
	}

	var newCart models.Cart
	tx := database.DB.Model(models.Cart{}).Where("goods_id = ?", intGoodsID).Where("product_id = ?", intProductID).First(&newCart)
	if tx.Error != nil {
		var goodsSpecifications []GoodsSpecification
		if productData.GoodsSpecificationIds != "" {
			goodsSpecificationIDs := strings.Split(productData.GoodsSpecificationIds, "_")

			database.DB.Raw("select gs.* s.name from goods_specification gs "+
				"inner json specification s on s.id = gs.specification_id where gs.goods_id = ? and gs.id in ?",
				intGoodsID, strings.Join(goodsSpecificationIDs, ",")).Scan(&goodsSpecifications)
		}

		goodsSpecificationJson, _ := json.Marshal(goodsSpecifications)
		database.DB.Model(models.Cart{}).Where("id = ?", intID).Updates(map[string]interface{}{
			"number":                         intNumber,
			"goods_specification_name_value": goodsSpecificationJson,
			"retail_price":                   productData.RetailPrice,
			"market_price":                   productData.RetailPrice,
			"product_id":                     intProductID,
			"goods_sn":                       productData.GoodsSn,
		})
	} else {
		newNumber := intNumber + newCart.Number
		if db.Error != nil || productData.GoodsNumber < newNumber {
			c.JSON(http.StatusBadRequest, "库存不足")
		}
		database.DB.Where("id = ?", newCart.ID).Delete(&models.Cart{})

		database.DB.Model(models.Cart{}).Where("id = ?", intID).Updates(map[string]interface{}{
			"number":                         newNumber,
			"goods_specification_name_value": newCart.GoodsSpecificationNameValue,
			"goods_specification_ids":        newCart.GoodsSpecificationIds,

			"retail_price":                   newCart.RetailPrice,
			"market_price":                   newCart.RetailPrice,
			"product_id":                     intProductID,
			"goods_sn":                       productData.GoodsSn,
		})
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))

}
