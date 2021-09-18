package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

type CommentCountRtnJson struct {
	AllCount    int64 `json:"all_count"`
	HasPicCount int `json:"has_pic_count"`
}

// Count 评论数量
func Count(c *gin.Context) {
	typeId := c.DefaultQuery("typeId", "0")
	valueId := c.DefaultQuery("valueId", "0")

	intTypeId, _ := strconv.Atoi(typeId)
	intValueId, _ := strconv.Atoi(valueId)

	var Total int64
	database.DB.Model(models.Comment{}).Where("type_id = ?", intTypeId).Where("value_id = ?", intValueId).Count(&Total)

	var list []models.Comment
	database.DB.Raw("select c.* "+
		"from comment c "+
		"right join comment_picture cp "+
		"on c.id = cp.comment_id "+
		"where c.type_id = ? and c.value_id = ?", typeId, valueId).Scan(&list)

	hasPicCount := len(list)
	c.JSON(http.StatusOK, utils.SuccessReturn(CommentCountRtnJson{
		AllCount:    Total,
		HasPicCount: hasPicCount,
	}))
}
