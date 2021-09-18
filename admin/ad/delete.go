package adminAd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shop_mater/database"
	"shop_mater/models"
)

type DeleteInput struct {
	ID int `json:"id"`
}

// Delete 广告删除
func Delete(c *gin.Context) {
	var InputData DeleteInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println(err)
		return
	}

	database.DB.Where("id = ?", InputData.ID).Delete(&models.Ad{})
}
