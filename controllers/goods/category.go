package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// CategoryRtnJson 商品分类  返回数据
type CategoryRtnJson struct {
	Category        models.Category   `json:"category"`
	ParentCategory  models.Category   `json:"parent_category"`
	BrotherCategory []models.Category `json:"brother_category"`
}

// Category 商品分类
func Category(c *gin.Context) {

	goodsId := c.DefaultQuery("id", "0")
	intGoodsId, _ := strconv.Atoi(goodsId)

	var category models.Category
	database.DB.Model(models.Category{}).Where("id = ?", intGoodsId).First(&category)
	var parentCategory models.Category
	database.DB.Model(models.Category{}).Where("id = ?", category.ParentId).First(&parentCategory)
	var brotherCategory []models.Category
	database.DB.Model(models.Category{}).Where("parent_id = ?", category.ParentId).Find(&brotherCategory)

	c.JSON(http.StatusOK, utils.SuccessReturn(CategoryRtnJson{
		Category:        category,
		ParentCategory:  parentCategory,
		BrotherCategory: brotherCategory,
	}))
}
