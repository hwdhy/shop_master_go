package category

import (
	"shop_mater/database"
	"shop_mater/models"
)

// GetChildCategoryId 获取分类下子分类ID集合
func GetChildCategoryId(categoryID int) []int64 {
	var childCategory []models.Category
	database.DB.Model(models.Category{}).Where("parent_id = ?", categoryID).Limit(10000).Find(&childCategory)

	var childCategoryIDs []int64
	for _, item := range childCategory {
		childCategoryIDs = append(childCategoryIDs, item.ID)
	}
	return childCategoryIDs
}

// GetCategoryWhereIn 获取当前分类和子分类ID集合
func GetCategoryWhereIn(categoryID int) []int64 {

	childCategoryId := GetChildCategoryId(categoryID)
	childCategoryId = append(childCategoryId, int64(categoryID))

	return childCategoryId
}
