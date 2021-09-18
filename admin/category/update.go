package adminCategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type CategoryUpdateInput struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ParentID     int    `json:"parent_id"`
	SortOrder    int    `json:"sort_order"`
	WapBannerUrl string `json:"wap_banner_url"`
	FrontName    string `json:"front_name"`
	FrontDesc    string `json:"front_desc"`
}

// Update 分类修改
func Update(c *gin.Context) {

	var InputData CategoryUpdateInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	var categoryData models.Category
	categoryData.Name = InputData.Name
	categoryData.ParentId = InputData.ParentID
	categoryData.SortOrder = InputData.SortOrder
	categoryData.WapBannerUrl = InputData.WapBannerUrl
	categoryData.FrontName = InputData.FrontName
	categoryData.FrontDesc = InputData.FrontDesc

	if db := database.DB.Model(models.Category{}).Where("id = ?", InputData.ID).Updates(&categoryData); db.Error != nil {
		c.JSON(http.StatusBadRequest, "")
	}
}
