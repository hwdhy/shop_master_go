package catalog

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type CatelogReturn struct {
	CurrentCategory CurCategory `json:"current_category"`
}

// Current 分类ID查询及其子分类集合
func Current(c *gin.Context) {
	categoryID := c.DefaultQuery("id", "0")

	var currentCategory models.Category
	if categoryID != "0" {
		database.DB.Model(models.Category{}).Where("id = ?", categoryID).First(&currentCategory)
	}

	var catelogReturn CurCategory
	if currentCategory.ID != 0 {
		var subCategory []models.Category
		database.DB.Model(models.Category{}).Where("parent_id = ?", currentCategory.ID).Find(&subCategory)
		catelogReturn.Category = currentCategory
		catelogReturn.SubCategoryList = subCategory
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(CatelogReturn{
		CurrentCategory: catelogReturn,
	}))

}
