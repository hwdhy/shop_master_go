package goods

import (
	"shop_mater/database"
	"shop_mater/models"
)

// ProductDetail 产品详情
func ProductDetail(goodsId int) []models.Product {

	var ProductData []models.Product
	database.DB.Model(models.Product{}).Where("goods_id = ?", goodsId).Find(&ProductData)

	return ProductData
}
