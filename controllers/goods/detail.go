package goods

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/collect"
	"shop_mater/service/footprint"
	"shop_mater/service/goods"
	"shop_mater/service/specification"
	"shop_mater/utils"
	"strconv"
)

// Attribute 属性
type Attribute struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// Comment 评论集合
type Comment struct {
	Count int64       `json:"count"`
	Data  CommentInfo `json:"data"`
}

// CommentInfo 评论信息
type CommentInfo struct {
	Content  string                  `json:"content"`
	AddTime  int64                   `json:"add_time"`
	NickName string                  `json:"nick_name"`
	Avatar   string                  `json:"avatar"`
	PicList  []models.CommentPicture `json:"pic_list"`
}

// DetailRtnJson 详情返回数据
type DetailRtnJson struct {
	SkuRtnJson
	Goods          models.Goods        `json:"info"`
	Galleries      []models.Gallery    `json:"gallery"`
	Attribute      []Attribute         `json:"attribute"`
	Issues         []models.GoodsIssue `json:"issue"`
	UserHasCollect int                 `json:"user_has_collect"`
	Comment        Comment             `json:"comment"`
	Brand          models.Brand        `json:"brand"`
}

type SkuRtnJson struct {
	ProductList       []models.Product                  `json:"product_list"`
	SpecificationList []specification.SpecificationItem `json:"specification_list"`
}

// Detail 产品详情
func Detail(c *gin.Context) {

	goodsId := c.DefaultQuery("id", "0")
	intGoodsId, _ := strconv.Atoi(goodsId)

	var goodsData models.Goods
	database.DB.Model(models.Goods{}).Where("id = ?", intGoodsId).First(&goodsData)

	//获取该产品下视图
	var gallery []models.Gallery
	database.DB.Model(models.Gallery{}).Where("goods_id = ?", intGoodsId).Limit(4).Find(&gallery)

	//获取商品属性
	var attributes []Attribute
	database.DB.Raw("select ga.value,a.name from goods_attribute ga inner join attribute a on ga.attribute_id = a.id where ga.goods_id = ?", goodsId).Scan(&attributes)

	//查询商品问题
	var issues []models.GoodsIssue
	database.DB.Model(models.GoodsIssue{}).Find(&issues)

	//品牌信息
	var brandData models.Brand
	database.DB.Model(models.Brand{}).Where("id = ?", goodsData.BrandID).First(&brandData)

	//评论
	var commentCount int64
	database.DB.Model(models.Comment{}).Where("value_id = ?", intGoodsId).Where("type_id = ?", 0).Count(&commentCount)
	var hotComment models.Comment
	database.DB.Model(models.Comment{}).Where("value_id = ?", intGoodsId).Where("type_id = ?", 0).First(&hotComment)

	var commentInfo CommentInfo

	if hotComment.ID != 0 {
		var commentUser models.User
		database.DB.Model(models.User{}).Select("nickname", "username", "avatar").Where("id = ?", hotComment.UserID).Find(&commentUser)

		content, _ := base64.StdEncoding.DecodeString(hotComment.Content)

		var commentPicture []models.CommentPicture
		database.DB.Model(models.CommentPicture{}).Where("comment_id = ?", hotComment.ID).Find(&commentPicture)

		commentInfo = CommentInfo{
			Content:  string(content),
			AddTime:  hotComment.AddTime,
			NickName: commentUser.Nickname,
			Avatar:   commentUser.Avatar,
			PicList:  commentPicture,
		}
	}

	commentVal := Comment{
		Count: commentCount,
		Data:  commentInfo,
	}

	loginUserID := base.GetLoginUserID()

	userHasCollect := collect.IsUserHasCollect(loginUserID, 0, intGoodsId)

	//加入浏览记录
	footprint.AddFootprint(int64(loginUserID), int64(intGoodsId))

	plist := goods.ProductDetail(intGoodsId)
	slist := specification.List(intGoodsId)

	c.JSON(http.StatusOK, utils.SuccessReturn(DetailRtnJson{
		SkuRtnJson: SkuRtnJson{
			ProductList:       plist,
			SpecificationList: slist,
		},
		Goods:          goodsData,
		Galleries:      gallery,
		Attribute:      attributes,
		Issues:         issues,
		UserHasCollect: userHasCollect,
		Comment:        commentVal,
		Brand:          brandData,
	}))
}
