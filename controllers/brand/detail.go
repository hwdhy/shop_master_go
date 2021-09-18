package brand

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

type BrandDetailRtnJson struct {
	Data models.Brand `json:"data"`
}

// Detail 品牌详情
func Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))

	var brandData models.Brand
	database.DB.Model(models.Brand{}).Where("id = ?", id).First(&brandData)

	c.JSON(http.StatusOK, utils.SuccessReturn(BrandDetailRtnJson{
		Data: brandData,
	}))

}
