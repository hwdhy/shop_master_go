package footprint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type RequestBody struct {
	FootprintId int `json:"footprintId"`
}

// Delete 浏览记录删除
func Delete(c *gin.Context) {

	var requestBody RequestBody
	err := c.Bind(&requestBody)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	loginUserId := base.GetLoginUserID()

	var footprint models.Footprint
	database.DB.Model(models.Footprint{}).Debug().Where("id = ?", requestBody.FootprintId).Where("user_id = ?", loginUserId).First(&footprint)

	database.DB.Debug().Where("user_id = ?", loginUserId).Where("goods_id = ?", footprint.GoodsID).Delete(&models.Footprint{})

	c.JSON(http.StatusOK,utils.SuccessReturn("删除成功"))
}
