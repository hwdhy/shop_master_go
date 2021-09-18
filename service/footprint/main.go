package footprint

import (
	"shop_mater/database"
	"shop_mater/models"
	"time"
)

func AddFootprint(userId, GoodsId int64) {
	if userId > 0 && GoodsId > 0 {
		var footprint = models.Footprint{
			GoodsID: GoodsId,
			UserID:  userId,
			AddTime: time.Now().Unix(),
		}

		database.DB.Model(models.Footprint{}).Create(&footprint)
	}
}
