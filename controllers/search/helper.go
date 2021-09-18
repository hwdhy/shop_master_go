package search

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

func Helper(c *gin.Context) {

	keyword := c.DefaultQuery("keyword", "0")

	var resKeywords []models.Keywords
	database.DB.Model(models.Keywords{}).Debug().Select("keyword").Where("keyword like ?", "%"+keyword+"%").Distinct().Limit(10).Find(&resKeywords)

	c.JSON(http.StatusOK, utils.SuccessReturn(resKeywords))

}
