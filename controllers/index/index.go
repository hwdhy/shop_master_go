package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

//分类商品集合
type newCategoryGoods struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	GoodsList []models.Goods `json:"goods_list"`
}

// IndexReturnData 返回数据
type IndexReturnData struct {
	Banners      []models.Ad        `json:"banners"`
	Channels     []models.Channel   `json:"channels"`
	NewGoods     []models.Goods     `json:"new_goods"`
	HotGoods     []models.Goods     `json:"hot_goods"`
	BrandList    []models.Brand     `json:"brand_list"`
	TopicList    []models.Topic     `json:"topic_list"`
	CategoryList []newCategoryGoods `json:"category_list"`
}

// IndexIndex 首页数据
func IndexIndex(c *gin.Context) {
	//广告
	var banners []models.Ad
	database.DB.Model(models.Ad{}).Where("ad_position_id = ?", 1).Find(&banners)
	//渠道
	var channels []models.Channel
	database.DB.Model(models.Channel{}).Order("sort_order").Find(&channels)
	//最新商品
	var newGoods []models.Goods
	database.DB.Model(models.Goods{}).
		Select("id", "name", "list_pic_url", "retail_price").
		Where("is_new = ?", 1).Limit(4).Find(&newGoods)

	//最热商品
	var hotGoods []models.Goods
	database.DB.Model(models.Goods{}).
		Select("id", "name", "list_pic_url", "retail_price", "goods_brief").
		Where("is_hot = ?", 1).
		Limit(3).Find(&hotGoods)
	//品牌
	var brandList []models.Brand
	database.DB.Model(models.Brand{}).Where("is_new = ?", 1).Order("new_sort_order").Limit(4).Find(&brandList)
	//专题
	var topicList []models.Topic
	database.DB.Model(models.Topic{}).Limit(3).Find(&topicList)
	//一级分类
	var categoryList []models.Category
	database.DB.Model(models.Category{}).Where("parent_id = ?", 0).Find(&categoryList)

	//查询二级分类下的商品
	var newList []newCategoryGoods
	for _, item := range categoryList {
		//二级分类集合
		var categoryIDs []models.Category
		database.DB.Model(models.Category{}).Select("id").Where("parent_id = ?", item.ID).Find(&categoryIDs)
		//取出二级分类ID
		var MapIDs []int64
		for _, v := range categoryIDs {
			MapIDs = append(MapIDs, int64(v.ID))
		}
		//二级分类ID查询商品
		var categoryGoods []models.Goods
		database.DB.Select("id", "name", "list_pic_url", "retail_price").Where("category_id in (?)", MapIDs).Limit(7).Find(&categoryGoods)

		newList = append(newList, newCategoryGoods{
			ID:        int(item.ID),
			Name:      item.Name,
			GoodsList: categoryGoods,
		})
	}
	//返回数据
	indexReturnData := IndexReturnData{
		Banners:      banners,
		Channels:     channels,
		NewGoods:     newGoods,
		HotGoods:     hotGoods,
		BrandList:    brandList,
		TopicList:    topicList,
		CategoryList: newList,
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(indexReturnData))
}
