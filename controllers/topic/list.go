package topic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// List 专题列表
func List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	var topicData []models.Topic
	var Total int64

	var current = (page - 1) * size
	database.DB.Debug().Model(models.Topic{}).
		Count(&Total).
		Select("id", "title", "price_info", "scene_pic_url", "subtitle").
		Offset(current).Limit(size).Find(&topicData)

	c.JSON(http.StatusOK, utils.SuccessReturn(
		utils.PageData{
			Count:     int(Total),
			Current:   current,
			TotalPage: (int(Total) + size - 1) / size,
			PageSize:  size,
			Data:      topicData,
		},
	))
}
