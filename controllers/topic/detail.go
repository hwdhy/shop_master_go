package topic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// Detail 专题详情
func Detail(c *gin.Context) {

	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))

	var topic models.Topic

	database.DB.Model(models.Topic{}).Where("id = ?", id).First(&topic)
	c.JSON(http.StatusOK, utils.SuccessReturn(topic))
}
