package brand

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
)

// PageData 查询数据封装
type PageData struct {
	Count     int         `json:"count"`      //总数
	Current   int         `json:"current"`    //页数
	TotalPage int         `json:"total_page"` //总页数
	PageSize  int         `json:"page_size"`  //每页返回条数
	Data      interface{} `json:"data"`       //数据集合
}

// List 品牌列表
func List(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	var brandData []models.Brand
	var Total int64
	database.DB.Model(models.Brand{}).Select("id", "name", "floor_price", "app_list_pic_url").Count(&Total).Offset(page).Limit(size).Find(&brandData)

	pageData := PageData{
		Count:     int(Total),
		Current:   page,
		TotalPage: (int(Total) + size - 1) / size,
		PageSize:  size,
		Data:      brandData,
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(pageData))
}
