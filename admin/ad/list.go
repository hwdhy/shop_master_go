package adminAd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
)

// AdListInput 广告列表输入参数
type AdListInput struct {
	Current  int `json:"current"`
	PageSize int `json:"page_size"`
}

type AdListOutput struct {
	Current  int                `json:"current"`
	PageSize int                `json:"page_size"`
	Total    int64              `json:"total"`
	Data     []AdListOutputData `json:"data"`
}

type AdListOutputData struct {
	ImageUrl string `json:"image_url"`
	Link     string `json:"link"`
	Content  string `json:"content"`
	Enabled  int    `json:"enabled"`
	Name     string `json:"name"`
	ID       int    `json:"id"`
}

// List 广告列表
func List(c *gin.Context) {

	var InputData AdListInput
	err := c.Bind(&InputData)
	if err != nil {
		fmt.Println("bind:", err)
		return
	}

	var Total int64
	Offset := (InputData.Current - 1) * InputData.PageSize

	var Output []AdListOutputData
	Total = database.DB.Debug().Raw("select a.id,a.image_url,a.link,a.content,a.enabled,ap.name from ad a "+
		"left join ad_position ap on a.ad_position_id = ap.id where a.deleted_at is null offset ? limit ?  ", Offset, InputData.PageSize).Scan(&Output).RowsAffected

	c.JSON(http.StatusOK, AdListOutput{
		Current:  InputData.Current,
		PageSize: InputData.PageSize,
		Total:    Total,
		Data:     Output,
	})
}
