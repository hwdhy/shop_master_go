package address

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/region"
	"shop_mater/utils"
	"strconv"
)

// Detail 收货地址详情
func Detail(c *gin.Context) {

	id := c.DefaultQuery("id", "0")

	intID, _ := strconv.Atoi(id)

	var address models.Address
	db := database.DB.Model(models.Address{}).Where("id = ?", intID).Where("user_id = ?", base.GetLoginUserID()).First(&address)
	if db.Error != nil {
		fmt.Println("not found!")
		return
	}

	provinceName := region.GetRegionName(address.ProvinceID)
	CityName := region.GetRegionName(address.CityID)
	districtName := region.GetRegionName(address.DistrictID)

	val := AddressListRtnJson{
		Address:      address,
		ProvinceName: provinceName,
		CityName:     CityName,
		DistrictName: districtName,
		FullRegion:   provinceName + CityName + districtName,
	}
	c.JSON(http.StatusOK, utils.SuccessReturn(val))

}
