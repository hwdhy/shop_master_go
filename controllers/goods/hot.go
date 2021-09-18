package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/utils"
)

// Banner 横幅
type Banner struct {
	Url    string `json:"url"`
	Name   string `json:"name"`
	ImgUrl string `json:"img_url"`
}

type HotRtnJson struct {
	BannerInfo Banner `json:"banner_info"`
}

// Hot 热门推荐
func Hot(c *gin.Context) {
	banner := Banner{
		Url:    "",
		Name:   "大家都在买的好物",
		ImgUrl: "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png",
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(HotRtnJson{
		BannerInfo: banner,
	}))
}
