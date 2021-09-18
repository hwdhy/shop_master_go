package region

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// Info 地区信息
func Info(c *gin.Context) {

	regionId := c.DefaultQuery("regionId", "0")
	intRegionId, _ := strconv.Atoi(regionId)

	var region models.Region

	database.DB.Model(models.Region{}).Where("id = ?", intRegionId).First(&region)

	c.JSON(http.StatusOK, utils.SuccessReturn(region))
}
