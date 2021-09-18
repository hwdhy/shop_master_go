package specification

import (
	"shop_mater/database"
	"shop_mater/models"
)

// SpecificationData 规格数据
type SpecificationData struct {
	models.GoodsSpecification
	Name string
}

type SpecificationItem struct {
	SpecificationId int                 `json:"specification_id"` //规格ID
	Name            string              `json:"name"`             //名称
	List            []SpecificationData `json:"list"`             //规格数据
}

// List 规格列表
func List(goodsId int) []SpecificationItem {
	var specificationData []SpecificationData

	database.DB.Raw("select gs.*,s.name from goods_specification gs inner join "+
		"specification s on gs.specification_id=s.id where gs.specification_id=?", goodsId).Scan(&specificationData)

	label := make(map[int]int)
	specificationList := make([]SpecificationItem, 0)
	var idx = 0

	for _, item := range specificationData {

		if v, ok := label[item.ID]; ok {
			specificationList[v].List = append(specificationList[v].List, item)
		} else {
			specificationList = append(specificationList, SpecificationItem{
				SpecificationId: item.ID,
				Name:            item.Name,
				List:            []SpecificationData{item},
			})

			label[item.ID] = idx
			idx += 1

		}
	}
	return specificationList
}
