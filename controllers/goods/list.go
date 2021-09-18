package goods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/service/category"
	"shop_mater/utils"
	"strconv"
	"strings"
	"time"
)

// FilterCategory 分类
type FilterCategory struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

// ListRtnJson 产品列表返回数据
type ListRtnJson struct {
	PageData
	FilterCategory []FilterCategory `json:"filter_category"`
	GoodsList      []models.Goods   `json:"goods_list"`
}

// PageData 查询数据封装
type PageData struct {
	Count     int         `json:"count"`      //总数
	TotalPage int         `json:"total_page"` //总页数
	Current   int         `json:"current"`    //当前页数
	PageSize  int         `json:"page_size"`  //每页返回条数
	Data      interface{} `json:"data"`       //数据集合
}

// List 商品列表
func List(c *gin.Context) {
	categoryId := c.DefaultQuery("categoryId", "0")
	brandId := c.DefaultQuery("brandId", "0")
	keyword := c.DefaultQuery("keyword", "0")
	isNew := c.DefaultQuery("isNew", "0")
	isHot := c.DefaultQuery("isHot", "0")
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")
	sort := c.DefaultQuery("sort", "0")
	order := c.DefaultQuery("order", "0")

	db := database.DB.Model(models.Goods{})

	intSize, _ := strconv.Atoi(size)
	if intSize == 0 {
		intSize = 20
	}
	intPage, _ := strconv.Atoi(page)
	keyword = strings.TrimSpace(keyword)
	if keyword != "0" {
		//TODO 将关键词加入搜索历史记录
		var searchHistory models.SearchHistory
		searchHistory.AddTime = time.Now().Unix()
		searchHistory.UserID = base.GetLoginUserID()
		searchHistory.Keyword = keyword
		database.DB.Model(models.SearchHistory{}).Create(&searchHistory)

		db = db.Where("name like ?", "%"+keyword+"%")
	}

	if isNew != "0" {
		db = db.Where("is_new = ?", isNew)
	}

	if isHot != "0" {
		db = db.Where("is_hot = ?", isHot)
	}

	if brandId != "0" {
		db = db.Where("brand_id = ?", brandId)
	}

	var categoryIDs []models.Goods
	db.Select("category_id").Limit(10000).Find(&categoryIDs)

	var categoryIdsMap []int64
	for _, item := range categoryIDs {
		categoryIdsMap = append(categoryIdsMap, item.CategoryId)
	}
	categoryIdsMap = utils.SliceRepeat(categoryIdsMap)

	var filterCategories = []FilterCategory{
		{Id: 0, Name: "全部", Checked: false},
	}

	//通过分类ID查询父ID
	if len(categoryIdsMap) > 0 {
		var parentCategory []models.Category
		database.DB.Model(models.Category{}).Select("parent_id").Where("id in (?)", categoryIdsMap).Limit(10000).Find(&parentCategory)

		var parentCategoryIDsMap []int64
		for _, item := range parentCategory {
			parentCategoryIDsMap = append(parentCategoryIDsMap, int64(item.ParentId))
		}
		parentCategoryIDsMap = utils.SliceRepeat(parentCategoryIDsMap)

		//通过父ID查询所有分类信息
		var parentCategoryData []models.Category
		database.DB.Model(models.Category{}).
			Where("id in (?)", parentCategoryIDsMap).
			Select("id", "name").Order("sort_order").Find(&parentCategoryData)

		for _, item := range parentCategoryData {
			id := item.ID
			name := item.Name
			checked := categoryId == "0" && id == 0

			filterCategories = append(filterCategories, FilterCategory{
				Id:      id,
				Name:    name,
				Checked: checked,
			})
		}
	}

	if categoryId != "0" {
		categoryIntId, _ := strconv.Atoi(categoryId)
		if categoryIntId > 0 {
			db = db.Where("category_id in (?)", category.GetCategoryWhereIn(categoryIntId))
		}
	}

	if sort == "price" {
		orderStr := "retail_price asc"
		if order == "desc" {
			orderStr = orderStr + " desc"
		}
		db = db.Order(orderStr)

	} else {
		db = db.Order("id desc")
	}

	var rawData []models.Goods
	Offset := (intPage - 1) * intSize
	var Total int64
	db.Debug().Select("id", "name", "list_pic_url", "retail_price").Count(&Total).Offset(Offset).Limit(intSize).Find(&rawData)

	c.JSON(http.StatusOK, utils.SuccessReturn(
		ListRtnJson{
			FilterCategory: filterCategories,
			GoodsList:      rawData,
			PageData: PageData{
				Current:   intPage,
				PageSize:  intSize,
				Count:     int(Total),
				TotalPage: (int(Total) + intSize - 1) / intSize,
				Data:      rawData,
			},
		}))
}
