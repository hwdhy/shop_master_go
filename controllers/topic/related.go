package topic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

// Related 专题关联
func Related(c *gin.Context) {

	var topicData []models.Topic

	database.DB.Model(models.Topic{}).
		Select("id", "title", "price_info", "scene_pic_url", "subtitle").
		Limit(4).Find(&topicData)

	c.JSON(http.StatusOK, utils.SuccessReturn(topicData))
}
