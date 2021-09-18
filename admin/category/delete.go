package adminCategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type CategoryDeleteInput struct {
	ID int64 `json:"id"` //ID
}

// Delete 分类删除
func Delete(c *gin.Context) {
	var InputData CategoryDeleteInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println(err)
		return
	}

	var repeatData models.Category
	//判断是否一级分类    一级分类不让删除
	database.DB.Model(models.Category{}).Select("parent_id").Where("id = ?").First(&repeatData)

	if repeatData.ParentId == 0 {
		c.JSON(http.StatusBadRequest, "一级分类不能删除")
		return
	}

	database.DB.Where("id = ?", InputData.ID).Delete(&models.Category{})
}
