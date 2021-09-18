package cart

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/region"
	"shop_mater/utils"
	"strconv"
)

type CheckoutRtnJson struct {
	Address          CartAddress         `json:"checkedAddress"`
	FreightPrice     float64             `json:"freightPrice"`
	CheckedCoupon    []models.UserCoupon `json:"checkedCoupon"`
	CouponList       []models.UserCoupon `json:"couponList"`
	CouponPrice      float64             `json:"couponPrice"`
	CheckedGoodsList []models.Cart       `json:"checkedGoodsList"`
	GoodsTotalPrice  float64             `json:"goodsTotalPrice"`
	OrderTotalPrice  float64             `json:"orderTotalPrice"`
	ActualPrice      float64             `json:"actualPrice"`
}

type CartAddress struct {
	models.Address
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

// CheckOut 确认购物车信息
func CheckOut(c *gin.Context) {
	addressId := c.DefaultQuery("addressId", "0")
	intAddressId, _ := strconv.Atoi(addressId)

	var addressData models.Address
	var db *gorm.DB
	if addressId != "0" {
		db = database.DB.Debug().Model(models.Address{}).Where("is_default = 1").Where("user_id = ?", base.GetLoginUserID()).First(&addressData)
	} else {
		db = database.DB.Debug().Model(models.Address{}).Where("id = ?", intAddressId).Where("user_id = ?", base.GetLoginUserID()).First(&addressData)
	}

	var customAddress CartAddress

	if db.Error == nil {
		customAddress.Address = addressData
		customAddress.ProvinceName = region.GetRegionName(addressData.ProvinceID)
		customAddress.CityName = region.GetRegionName(addressData.CityID)
		customAddress.DistrictName = region.GetRegionName(addressData.DistrictID)
		customAddress.FullRegion = customAddress.ProvinceName + customAddress.CityName + customAddress.DistrictName
	}
	var freightPrice = 0.0

	cartData := GetCart()
	var checkedGoodsList []models.Cart
	for _, val := range cartData.CartList {
		if val.Checked == 1 {
			checkedGoodsList = append(checkedGoodsList, val)
		}
	}

	var couponList []models.UserCoupon
	database.DB.Model(models.UserCoupon{}).Find(&couponList)

	var couponPrice = 0.0
	goodsTotalPrice := cartData.CartTotal.CheckedGoodsAmount
	orderTotalPrice := cartData.CartTotal.CheckedGoodsAmount + freightPrice - couponPrice

	actualPrice := orderTotalPrice - 0

	c.JSON(http.StatusOK, utils.SuccessReturn(CheckoutRtnJson{
		Address:          customAddress,
		FreightPrice:     freightPrice,
		CouponList:       couponList,
		CouponPrice:      couponPrice,
		CheckedGoodsList: checkedGoodsList,
		GoodsTotalPrice:  goodsTotalPrice,
		OrderTotalPrice:  orderTotalPrice,
		ActualPrice:      actualPrice,
	}))
}
