package adminAd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type AdUpdateInput struct {
	ID           int    `json:"id"`
	AdPositionId int    `json:"ad_position_id"`
	ImageUrl     string `json:"image_url"`
	Link         string `json:"link"`
	Content      string `json:"content"`
	Enabled      int8   `json:"enabled"`
}

// Update 广告修改
func Update(c *gin.Context) {

	var InputData AdUpdateInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}
	var adData models.Ad
	adData.AdPositionId = InputData.AdPositionId
	adData.ImageUrl = InputData.ImageUrl
	adData.Link = InputData.Link
	adData.Content = InputData.Content
	adData.Enabled = InputData.Enabled

	if db := database.DB.Model(models.Ad{}).Where("id = ?", InputData.ID).Updates(&adData); db.Error != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	c.JSON(http.StatusOK, "")

}
