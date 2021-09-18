package adminGoods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
)

type GoodsListInput struct {
	Current  int `json:"current"`   //当前页码
	PageSize int `json:"page_size"` //每页条数
}

type GoodsListOutput struct {
	Current  int            `json:"current"`   //当前页码
	PageSize int            `json:"page_size"` //每页条数
	Total    int64          `json:"total"`     //返回数据总和
	Data     []models.Goods `json:"data"`      //返回数据集合
}

// List 产品列表
func List(c *gin.Context) {

	var inputData GoodsListInput
	err := c.Bind(&inputData)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	Offset := (inputData.Current - 1) * inputData.PageSize
	var Total int64

	var goodsData []models.Goods
	database.DB.Debug().Model(models.Goods{}).Count(&Total).Order("updated_at desc,id desc").Offset(Offset).Limit(inputData.PageSize).Find(&goodsData)

	fmt.Println(goodsData[9].ID)

	c.JSON(http.StatusOK, GoodsListOutput{
		Current:  inputData.Current,
		PageSize: inputData.PageSize,
		Total:    Total,
		Data:     goodsData,
	})
}
