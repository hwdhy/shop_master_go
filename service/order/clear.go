package order

import (
	"shop_mater/database"
	"shop_mater/models"
)

// ClearBuyGoods 清除购买的商品
func ClearBuyGoods(userId int) {
	database.DB.Model(models.Cart{}).Where("user_id = ?", userId).Where("session_id = ?", "1").Where("checked = 1").Delete(&models.Cart{})

}
