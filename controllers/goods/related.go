package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// RelatedRtnJson 返回数据
type RelatedRtnJson struct {
	GoodsList []models.Goods `json:"goods_list"`
}

// Related 关联商品
func Related(c *gin.Context) {
	goodsID, _ := strconv.Atoi(c.DefaultQuery("id", "0"))

	var relatedGoods []models.GoodsRelated
	//查询关联商品ID
	database.DB.Model(models.GoodsRelated{}).Select("related_goods_id").Where("goods_id = ?", goodsID).Find(&relatedGoods)
	//关联商品ID集合
	var relatedGoodsIdMap []int64
	for _, item := range relatedGoods {
		relatedGoodsIdMap = append(relatedGoodsIdMap, item.RelatedGoodsID)
	}

	//关联商品数据
	var relatedGoodsData []models.Goods
	if len(relatedGoodsIdMap) == 0 {
		var goodsData models.Goods
		database.DB.Model(models.Goods{}).Where("id = ?", goodsID).First(&goodsData)

		database.DB.Model(models.Goods{}).Select("id", "name", "list_pic_url", "retail_price").Where("category_id = ?", goodsData.CategoryId).Limit(8).Find(&relatedGoodsData)
	} else {
		database.DB.Model(models.Goods{}).Select("id", "name", "list_pic_url", "retail_price").Where("id in (?)", relatedGoodsIdMap).Find(&relatedGoodsData)
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(RelatedRtnJson{
		GoodsList: relatedGoodsData,
	}))
}
