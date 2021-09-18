package search

import (
	"github.com/gin-gonic/gin"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
)

// Clear 清除历史记录
func Clear(c *gin.Context) {
	database.DB.Where("user_id = ?", base.GetLoginUserID()).Delete(&models.SearchHistory{})

	return
}
