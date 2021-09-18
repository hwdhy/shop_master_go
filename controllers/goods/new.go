package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/utils"
)

type NewRtnJson struct {
	BannerInfo Banner `json:"banner_info"`
}

// New 新品推荐
func New(c *gin.Context) {
	banner := Banner{
		Url:    "",
		Name:   "坚持初心，为你寻觅世间好物",
		ImgUrl: "http://yanxuan.nosdn.127.net/8976116db321744084774643a933c5ce.png",
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(NewRtnJson{
		BannerInfo: banner,
	}))
}
