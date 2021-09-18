package footprint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"sort"
	"time"
)

type FootprintListRtnJson struct {
	models.Footprint
	Name          string `json:"name"`
	ListPicUrl    string `json:"list_pic_url"`
	GoodsBrief    string `json:"goods_brief"`
	RetailPrice   int64  `json:"retail_price"`
	AddTimeString string `json:"add_time_string"`
}

// List 浏览记录
func List(c *gin.Context) {

	var listData []FootprintListRtnJson
	database.DB.Debug().Raw("select f.*,g.name,g.list_pic_url,g.goods_brief,g.retail_price "+
		"from footprint f "+
		"inner join goods g "+
		"on f.goods_id = g.id "+
		"where f.user_id = ? and f.deleted_at is null  order by id desc", base.GetLoginUserID()).Scan(&listData)

	var rvData []FootprintListRtnJson
	var goodsIDs []int

	for _, item := range listData {
		//去重
		if !utils.ContainsInt(goodsIDs, int(item.GoodsID)) {
			footprintTime := time.Unix(item.AddTime, 0)
			nowTime := time.Now()
			yesterdayTime := nowTime.Add(-24 * time.Hour)
			yesYesterdayTime := yesterdayTime.Add(-24 * time.Hour)

			if utils.DateEqual(footprintTime, nowTime) {
				item.AddTimeString = "今天"
			} else if utils.DateEqual(footprintTime, yesterdayTime) {
				item.AddTimeString = "昨天"
			} else if utils.DateEqual(footprintTime, yesYesterdayTime) {
				item.AddTimeString = "前天"
			}
			goodsIDs = append(goodsIDs, int(item.GoodsID))
			rvData = append(rvData, item)
		}
	}

	//todo 排序
	sort.Slice(rvData, func(i, j int) bool {
		return rvData[i].AddTime > rvData[j].AddTime
	})

	c.JSON(http.StatusOK, utils.SuccessReturn(rvData))

}
