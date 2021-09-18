package adminAdPosition

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type AdPositionOutput struct {
	Data []AdPositionOutputData `json:"data"`
}

type AdPositionOutputData struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

// List 广告位置列表
func List(c *gin.Context) {

	var adPositionData []models.AdPosition
	database.DB.Model(models.AdPosition{}).Find(&adPositionData)

	var resData AdPositionOutputData
	var resDatas []AdPositionOutputData

	for _, item := range adPositionData {
		resData.ID = int(item.ID)
		resData.Value = item.Name

		resDatas = append(resDatas, resData)
	}

	c.JSON(http.StatusOK, AdPositionOutput{
		Data: resDatas,
	})
}
