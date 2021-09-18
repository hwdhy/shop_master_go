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

type RequestBody struct {
	ID int `json:"id"`
}

// Delete 删除收货地址
func Delete(c *gin.Context) {

	var rb RequestBody
	err := c.Bind(&rb)
	if err != nil {
		fmt.Println("bind err:", err)
		return
	}

	database.DB.Where("id = ?", rb.ID).Where("user_id = ?", base.GetLoginUserID()).Delete(&models.Address{})

	c.JSON(http.StatusOK, utils.SuccessReturn("删除成功"))

}
