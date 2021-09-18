package region

import (
	"shop_mater/database"
	"shop_mater/models"
)

// GetRegionName 获取地区名称
func GetRegionName(regionID int) string {

	var regionData models.Region
	database.DB.Model(models.Region{}).Where("id = ?", regionID).First(&regionData)

	return regionData.Name
}
