package address

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type AddressSaveBody struct {
	Address    string `json:"address"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	IsDefault  bool   `json:"is_default"`
	Mobile     string `json:"mobile"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`
	AddressId  int    `json:"address_id"`
}

func Save(c *gin.Context) {

	var saveBody AddressSaveBody

	err := c.Bind(&saveBody)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	address := saveBody.Address
	name := saveBody.Name
	mobile := saveBody.Mobile
	provinceId := saveBody.ProvinceId
	cityID := saveBody.CityId
	distinctId := saveBody.DistrictId
	isDefault := saveBody.IsDefault
	addressID := saveBody.AddressId
	userID := base.GetLoginUserID()

	var intIsDefault int
	if isDefault {
		intIsDefault = 1
	} else {
		intIsDefault = 2
	}

	intCityID := cityID
	intProvinceID := provinceId
	intDistinctID := distinctId

	addressData := models.Address{
		Address:    address,
		CityID:     intCityID,
		DistrictID: intDistinctID,
		ProvinceID: intProvinceID,
		Name:       name,
		Mobile:     mobile,
		UserID:     userID,
		IsDefault:  intIsDefault,
	}

	var intID int64
	if addressID == 0 {
		database.DB.Model(models.Address{}).Create(&addressData)
		intID = addressData.ID
	} else {
		database.DB.Model(models.Address{}).Where("id = ?", intID).Where("user_id = ?", userID).Update("is_default", 0)
	}

	if isDefault {
		database.DB.Exec("update address set is_default = 0 where id != ? and user_id = ?", intID, userID)
	}

	var addressInfo models.Address
	database.DB.Model(models.Address{}).Where("id = ?", intID).First(&addressInfo)

	c.JSON(http.StatusOK, utils.SuccessReturn(addressInfo))
}
