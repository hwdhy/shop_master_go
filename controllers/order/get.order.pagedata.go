package order

import (
	"shop_mater/models"
	"shop_mater/utils"
)

func GetOrderPageData(rawData []models.Order, page int, size int) utils.PageData {

	count := len(rawData)
	totalPages := (count + size - 1) / size

	var pageData []models.Order

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pageData = append(pageData, rawData[idx])
	}

	return utils.PageData{
		Count:     count,
		Current:   page,
		TotalPage: totalPages,
		PageSize:  size,
		Data:      pageData,
	}

}
