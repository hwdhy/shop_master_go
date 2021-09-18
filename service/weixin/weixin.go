package weixin

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shop_mater/config"
	"shop_mater/utils"
)

// ResUserInfo 返回用户信息
type ResUserInfo struct {
	UserInfo      WXUserInfo `json:"userInfo"`
	RawData       string     `json:"rawData"`
	Signature     string     `json:"signature"`
	EncryptedData string     `json:"encryptedData"`
	IV            string     `json:"iv"`
}

// WXUserInfo 微信用户信息
type WXUserInfo struct {
	OpenID    string    `json:"openId,omitempty"`
	NickName  string    `json:"nickName"`
	AvatarUrl string    `json:"avatarUrl"`
	Gender    int       `json:"gender"`
	Country   string    `json:"country"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	UnionID   string    `json:"unionId,omitempty"`
	Language  string    `json:"language"`
	Watermark Watermark `json:"watermark,omitempty"`
}

// Watermark 水印
type Watermark struct {
	AppID     string `json:"appid"`
	TimeStamp int64  `json:"timestamp"`
}

// WXLoginResponse 微信登录返回数据
type WXLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// Login 登录
func Login(code string, fullUserInfo ResUserInfo) *WXUserInfo {

	secret := config.Secret
	appId := config.APPID

	fmt.Println("code:", code)
	fmt.Println("secret:", secret)
	fmt.Println("appId:", appId)
	client := &http.Client{}

	url := "https://api.weixin.qq.com/sns/jscode2session"

	resp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	q := resp.URL.Query()
	q.Add("grant_type", "authorization_code")
	q.Add("js_code", code)
	q.Add("secret", secret)
	q.Add("appid", appId)

	resp.URL.RawQuery = q.Encode()

	res, err := client.Do(resp)
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}
	defer res.Body.Close()

	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("readAll err:", err)
		return nil
	}
	var ress WXLoginResponse
	err = json.Unmarshal(all, &ress)
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}
	s := sha1.New()
	s.Write([]byte(fullUserInfo.RawData + ress.SessionKey))
	sha1 := s.Sum(nil)
	sha1Hash := hex.EncodeToString(sha1)

	if fullUserInfo.Signature != sha1Hash {
		return nil
	}

	userInfo := DecryptUserInfoData(ress.SessionKey, fullUserInfo.EncryptedData, fullUserInfo.IV)

	return userInfo
}

func DecryptUserInfoData(sessionKey string, encryptedData string, iv string) *WXUserInfo {
	sk, _ := base64.StdEncoding.DecodeString(sessionKey)
	ed, _ := base64.StdEncoding.DecodeString(encryptedData)
	i, _ := base64.StdEncoding.DecodeString(iv)

	decryptedData, err := utils.AesCBCDecrypt(ed, sk, i)
	if err != nil {
		return nil
	}

	var wxUserInfo WXUserInfo
	_ = json.Unmarshal(decryptedData, &wxUserInfo)

	return &wxUserInfo
}
