package comment

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
	"time"
)

// Post 发表评论
func Post(c *gin.Context) {
	typeId := c.DefaultQuery("typeId", "0")
	valueId := c.DefaultQuery("valueId", "0")
	content := c.DefaultQuery("content", "0")

	intTypeId, _ := strconv.Atoi(typeId)
	intValueId, _ := strconv.Atoi(valueId)

	var comment = models.Comment{
		AddTime: time.Now().Unix(),
		Content: base64.StdEncoding.EncodeToString([]byte(content)),
		TypeID:  int8(intTypeId),
		UserID:  base.GetLoginUserID(),
		ValueID: intValueId,
	}

	if db := database.DB.Model(models.Comment{}).Create(&comment); db.Error != nil {
		c.JSON(http.StatusBadRequest, utils.SuccessReturn("添加评论失败"))
	} else {
		c.JSON(http.StatusOK, utils.SuccessReturn("添加评论成功"))
	}
}
