package catalog

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

// CurCategory 分类 二级分类集合
type CurCategory struct {
	models.Category
	SubCategoryList []models.Category `json:"sub_category_list"`
}

// CateLogIndexRtnJson 输出集合
type CateLogIndexRtnJson struct {
	CategoryList    []models.Category `json:"category_list"`
	CurrentCategory CurCategory       `json:"current_category"`
}

// Index 分类首页
func Index(c *gin.Context) {
	categoryId := c.DefaultQuery("id", "0")

	var categoryDatas []models.Category
	database.DB.Model(models.Category{}).Where("parent_id = ?", categoryId).Limit(10).Find(&categoryDatas)

	var currentCategory models.Category
	if categoryId != "" {
		database.DB.Model(models.Category{}).Where("parent_id = ?", categoryId).First(&currentCategory)
	}

	if currentCategory.ID == 0 {
		currentCategory = categoryDatas[0]
	}

	var curCategory CurCategory
	if currentCategory.ID > 0 {
		var subCategoryDatas []models.Category
		database.DB.Model(models.Category{}).Where("parent_id = ?", currentCategory.ID).Find(&subCategoryDatas)
		curCategory.Category = currentCategory
		curCategory.SubCategoryList = subCategoryDatas
	}

	cateLogIndexRtnJson := CateLogIndexRtnJson{
		CategoryList:    categoryDatas,
		CurrentCategory: curCategory,
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(cateLogIndexRtnJson))
}
