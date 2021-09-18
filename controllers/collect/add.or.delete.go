package collect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"time"
)

type AddorDeleteRtnJson struct {
	HandleType string `json:"handle_type"`
}

type RequestBody struct {
	TypeId  int `json:"typeId"`
	ValueID int `json:"valueId"`
}

// AddOrDelete 收藏 取消收藏
func AddOrDelete(c *gin.Context) {

	var rb RequestBody
	err := c.Bind(&rb)
	if err != nil {
		fmt.Println("参数错误")
		return
	}

	var collect models.Collect
	var rvjson AddorDeleteRtnJson

	db := database.DB.Model(models.Collect{}).Debug().
		Where("type_id = ?", rb.TypeId).
		Where("value_id = ?", rb.ValueID).Where("user_id = ?", base.GetLoginUserID()).First(&collect)

	if db.Error != nil {
		database.DB.Model(models.Collect{}).Create(&models.Collect{
			TypeId:  rb.TypeId,
			ValueId: rb.ValueID,
			UserId:  base.GetLoginUserID(),
			AddTime: time.Now().Unix(),
		})
		rvjson = AddorDeleteRtnJson{
			HandleType: "add",
		}
	} else {
		database.DB.Unscoped().Where("id = ?", collect.ID).Delete(&models.Collect{})
		rvjson = AddorDeleteRtnJson{HandleType: "delete"}
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(rvjson))
}
