package weixin

import (
	"fmt"
	"github.com/objcoding/wxpay"
	"log"
	"shop_mater/config"
)

type PayInfo struct {
	OpenId         string
	Body           string
	OutTradeNo     string
	TotalFee       int64
	SpbillCreateIp string
}

type Params map[string]string

func CreateUnifiedOrder(payInfo PayInfo) (wxpay.Params, error) {

	log.Printf("%+v", payInfo)
	appid := config.APPID
	mchid := config.Mch_id
	apikey := config.Apikey
	//notifyurl := config.Notify_url

	account := wxpay.NewAccount(appid, mchid, apikey, false)
	client := wxpay.NewClient(account)

	fmt.Println("1111111111111111")
	params := make(wxpay.Params)
	params.SetString("body", payInfo.Body).
		SetString("out_trade_no", payInfo.OutTradeNo).
		SetInt64("total_fee", payInfo.TotalFee).
		SetString("spbill_create_id", payInfo.SpbillCreateIp).
		//SetString("notify_url", notifyurl).
		SetString("trade_type", "APP")

	fmt.Println("2222222222222222222")
	return client.UnifiedOrder(params)
}
