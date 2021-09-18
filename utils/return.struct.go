package utils

import (
	"crypto/aes"
	"crypto/cipher"
)

type HTTPData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// PageData 查询数据封装
type PageData struct {
	Count     int         `json:"count"`      //总数
	Current   int         `json:"current"`    //页数
	TotalPage int         `json:"total_page"` //总页数
	PageSize  int         `json:"page_size"`  //每页返回条数
	Data      interface{} `json:"data"`       //数据集合
}

// SuccessReturn 请求成功信息
func SuccessReturn(v interface{}) HTTPData {

	return HTTPData{
		ErrNo:  0,
		ErrMsg: "",
		Data:   v,
	}
}

// ErrReturn 错误返回
func ErrReturn(errno int, errmsg string) HTTPData {
	return HTTPData{
		ErrNo:  errno,
		ErrMsg: errmsg,
		Data:   nil,
	}
}


func AesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}

	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedData := make([]byte, len(encryptData))
	mode.CryptBlocks(decryptedData, encryptData)
	decryptedData = PKCS7UnPadding(decryptedData)
	return decryptedData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}