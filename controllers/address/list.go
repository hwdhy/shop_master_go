package address

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/region"
	"shop_mater/utils"
)

type AddressListRtnJson struct {
	models.Address
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

// List 收货地址列表
func List(c *gin.Context) {
	var address []models.Address
	database.DB.Model(models.Address{}).Where("user_id = ?", base.GetLoginUserID()).Find(&address)

	rtnAddress := make([]AddressListRtnJson, 0)

	for _, item := range address {
		provinceName := region.GetRegionName(item.ProvinceID)
		cityName := region.GetRegionName(item.CityID)
		distinctName := region.GetRegionName(item.DistrictID)
		rtnAddress = append(rtnAddress, AddressListRtnJson{
			Address:      item,
			ProvinceName: provinceName,
			CityName:     cityName,
			DistrictName: distinctName,
			FullRegion:   provinceName + cityName + distinctName,
		})
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(rtnAddress))

}
