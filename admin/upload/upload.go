package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"path"
	"strings"
)

// Upload 上传
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")

	ext := path.Ext(file.Filename)

	replace := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	err := c.SaveUploadedFile(file, "D:/tupian/"+replace+ext)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url": replace + ext,
	})
}
