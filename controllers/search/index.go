package search

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_mater/controllers/base"
	"shop_mater/database"
	"shop_mater/models"
	"shop_mater/utils"
)

type SearchIndexRtnJson struct {
	DefaultKeyword     models.Keywords   `json:"default_keyword"`
	HistoryKeywordList []string          `json:"history_keyword_list"`
	HotKeywordList     []models.Keywords `json:"hot_keyword_list"`
}

func Index(c *gin.Context) {

	var keyword models.Keywords
	database.DB.Model(models.Keywords{}).Where("is_default = 1").Limit(1).First(&keyword)

	var hotKeywords []models.Keywords
	database.DB.Model(models.Keywords{}).Distinct().
		Select("keyword", "is_hot").Limit(10).Find(&hotKeywords)

	var historyKeywords []models.SearchHistory
	database.DB.Model(models.SearchHistory{}).Select("keyword").
		Where("user_id = ?", base.GetLoginUserID()).Distinct().Limit(10).Find(&historyKeywords)

	var arrayKeywords []string
	for _, item := range historyKeywords {
		arrayKeywords = append(arrayKeywords, item.Keyword)
	}

	c.JSON(http.StatusOK, utils.SuccessReturn(SearchIndexRtnJson{
		DefaultKeyword:     keyword,
		HistoryKeywordList: arrayKeywords,
		HotKeywordList:     hotKeywords,
	}))
}
