package cart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
	"strings"
)

type CartDeleteBody struct {
	ProductIds string `json:"productIds"`
}

// Delete 删除购物车商品
func Delete(c *gin.Context) {

	var deleteBody CartDeleteBody
	err := c.Bind(&deleteBody)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("productIdArray:", deleteBody)
	intUserID := base.GetLoginUserID()

	productIdArray := strings.Split(deleteBody.ProductIds, ",")
	database.DB.Where("product_id in ?", productIdArray).Where("user_id = ?", intUserID).Delete(&models.Cart{})

	c.JSON(http.StatusOK, utils.SuccessReturn(GetCart()))
}
