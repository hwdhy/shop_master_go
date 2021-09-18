package cart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strconv"
	"strings"
)

type CartCheckedBody struct {
	IsChecked  int         `json:"isChecked"`
	ProductIds interface{} `json:"productIds"`
}

// Checked 选择 取消选择
func Checked(c *gin.Context) {
	var cb CartCheckedBody

	err := c.Bind(&cb)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	intIsChecked := cb.IsChecked
	if cb.ProductIds == "" {
		c.JSON(http.StatusOK, "删除出错")
		return
	}

	var productIDsArray []string

	switch val := cb.ProductIds.(type) {
	case float64:
		productIDsArray = append(productIDsArray, strconv.Itoa(int(val)))
	case string:
		productIDsArray = strings.Split(val, ",")
	default:

	}
	database.DB.Model(models.Cart{}).Where("product_id  in ?", productIDsArray).Update("checked", intIsChecked)

	c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))
}
