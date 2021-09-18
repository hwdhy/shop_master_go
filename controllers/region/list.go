package region

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// List 区域列表
func List(c *gin.Context) {

	parentId := c.DefaultQuery("parentId", "0")
	intParentId, _ := strconv.Atoi(parentId)

	var regions []models.Region
	database.DB.Model(models.Region{}).Where("parent_id = ?", intParentId).Find(&regions)

	c.JSON(http.StatusOK, utils.SuccessReturn(regions))
}
