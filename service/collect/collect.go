package collect

import (
	"shop_mater/database"
	"shop_mater/models"
)

// IsUserHasCollect 判断当前用户是否收藏
func IsUserHasCollect(userId, typeId, valueId int) int {

	var collectData models.Collect
	err := database.DB.Model(models.Collect{}).Where("type_id = ?", typeId).Where("value_id = ?", valueId).Where("user_id = ?", userId).First(&collectData)
	if err != nil {
		return 0
	} else {
		return 1
	}
}
