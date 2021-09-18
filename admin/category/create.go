package adminCategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type CategoryCreateInput struct {
	Name         string `json:"name"`
	ParentID     int    `json:"parent_id"`
	SortOrder    int    `json:"sort_order"`
	WapBannerUrl string `json:"wap_banner_url"`
	FrontName    string `json:"front_name"`
	FrontDesc    string `json:"front_desc"`
}

// Create 分类创建
func Create(c *gin.Context) {
	var InputData CategoryCreateInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var categoryData models.Category
	categoryData.Name = InputData.Name
	categoryData.ParentId = InputData.ParentID
	categoryData.SortOrder = InputData.SortOrder
	categoryData.WapBannerUrl = InputData.WapBannerUrl
	categoryData.FrontName = InputData.FrontName
	categoryData.FrontDesc = InputData.FrontDesc

	if db := database.DB.Model(models.Category{}).Create(&categoryData); db.Error != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	c.JSON(http.StatusOK, "")
}
