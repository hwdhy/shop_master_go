package base

import (
	"shop_mater/service/login"
	"strconv"
)

func GetLoginUserID() int {
	intUserID, _ := strconv.Atoi(login.LoginUserId)
	return intUserID
}
