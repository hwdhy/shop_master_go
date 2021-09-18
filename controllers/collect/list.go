package collect

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type CollectListRtnJson struct {
	models.Collect
	Name        string  `json:"name"`
	ListPicUrl  string  `json:"list_pic_url"`
	GoodsBrief  string  `json:"goods_brief"`
	RetailPrice float64 `json:"retail_price"`
}

// List 收藏列表
func List(c *gin.Context) {

	typeId := c.DefaultQuery("typeId", "0")

	var list []CollectListRtnJson

	database.DB.Debug().Raw("select c.*,g.name,g.list_pic_url,g.goods_brief,g.retail_price "+
		"from collect c "+
		"left join goods g "+
		"on c.value_id = g.id "+
		"where c.user_id = ? and c.type_id = ? and c.deleted_at is null", base.GetLoginUserID(), typeId).Scan(&list)

	c.JSON(http.StatusOK, utils.SuccessReturn(list))

}
