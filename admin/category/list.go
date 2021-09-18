package adminCategory

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type CategoryListInput struct {
	Current  int   `json:"current"`   //当前页码
	PageSize int   `json:"page_size"` //每页条数
	ParentId int64 `json:"parent_id"` //父ID
}

type CategoryListOutput struct {
	Current  int               `json:"current"`   //当前页码
	PageSize int               `json:"page_size"` //每页条数
	Total    int64             `json:"total"`     //返回数据总和
	Data     []models.Category `json:"data"`      //返回数据集合
}

// List 分类列表
func List(c *gin.Context) {

	var InputData CategoryListInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	db := database.DB.Model(models.Category{})
	Offset := (InputData.Current - 1) * InputData.PageSize

	var Total int64
	var categoryData []models.Category
	db.Where("parent_id = ?", InputData.ParentId).Count(&Total).Offset(Offset).Limit(InputData.PageSize).Find(&categoryData)

	c.JSON(http.StatusOK, CategoryListOutput{
		Current:  InputData.Current,
		PageSize: InputData.PageSize,
		Total:    Total,
		Data:     categoryData,
	})
}
