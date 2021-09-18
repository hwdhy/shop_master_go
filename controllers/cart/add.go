package cart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
	"strings"
)

type CartAddBody struct {
	GoodsId   int `json:"goodsId"`
	ProductId int `json:"productId"`
	Number    int `json:"number"`
}

// Add 添加商品到购物车
func Add(c *gin.Context) {
	var cartAddBody CartAddBody
	fmt.Printf("%+v", cartAddBody.ProductId)
	err := c.Bind(&cartAddBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	intGoodsID := cartAddBody.GoodsId
	intProductID := cartAddBody.ProductId
	intNumber := cartAddBody.Number
	intUserID := base.GetLoginUserID()

	var goods models.Goods
	db := database.DB.Model(models.Goods{}).Where("id = ?", intGoodsID).First(&goods)
	if db.Error != nil || goods.IsDelete == 1 {
		c.JSON(http.StatusBadRequest, utils.SuccessReturn("商品已下架"))
		return
	}

	var productData models.Product
	db = database.DB.Debug().Model(models.Product{}).
		Where("goods_id = ?", intGoodsID).Where("id = ?", intProductID).First(&productData)
	if db.Error != nil || productData.GoodsNumber < intNumber {
		c.JSON(http.StatusBadRequest, utils.SuccessReturn("库存不足"))
		return
	}

	var cartData models.Cart
	db = database.DB.Model(models.Cart{}).Where("goods_id = ?", intGoodsID).Where("product_id = ?", intProductID).Where("user_id = ?", intUserID).First(&cartData)

	if db.Error != nil {
		var goodsSpecificationData []models.GoodsSpecification
		if productData.GoodsSpecificationIds != "" {
			goodsSpecificationIDs := strings.Split(productData.GoodsSpecificationIds, "_")
			var goodsSpecificationIDsMap []int
			for _, item := range goodsSpecificationIDs {
				a, _ := strconv.Atoi(item)
				goodsSpecificationIDsMap = append(goodsSpecificationIDsMap, a)
			}
			database.DB.Model(models.GoodsSpecification{}).Select("value").
				Where("goods_id = ?", intGoodsID).Where("id in (?)", goodsSpecificationIDsMap).Find(&goodsSpecificationData)
		}

		var vals []string
		for _, item := range goodsSpecificationData {
			vals = append(vals, item.Value)
		}

		cartDatas := models.Cart{
			GoodsId:                     intGoodsID,
			ProductID:                   intProductID,
			GoodsSn:                     productData.GoodsSn,
			GoodsName:                   goods.Name,
			ListPicUrl:                  goods.ListPicUrl,
			Number:                      intNumber,
			SessionID:                   "1",
			UserID:                      intUserID,
			RetailPrice:                 productData.RetailPrice,
			MarketPrice:                 productData.RetailPrice,
			GoodsSpecificationNameValue: strings.Join(vals, ";"),
			GoodsSpecificationIds:       productData.GoodsSpecificationIds,
			Checked:                     1,
		}

		database.DB.Model(models.Cart{}).Create(&cartDatas)

		c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))
	} else {
		if productData.GoodsNumber < (intNumber + cartData.Number) {
			c.JSON(http.StatusOK, utils.SuccessReturn("库存不足"))
			return
		}
		database.DB.Exec("update cart set number = number + ? where id = ? and goods_id = ? and product_id = ?", intNumber, cartData.ID, intGoodsID, intProductID)

		c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))
	}
}
