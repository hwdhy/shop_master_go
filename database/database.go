package database

import (
	"fmt"
	"gorm.io/gorm"
	"shop_mater/models"
)
import "gorm.io/driver/postgres"

var DB *gorm.DB

// InitDatabase 连接初始化数据库
func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=dhyy dbname=shop port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	RegisterTable(db)

	DB = db
}

// RegisterTable 注册表
func RegisterTable(db *gorm.DB) {
	_ = db.AutoMigrate(
		models.Ad{},
		models.AdPosition{},
		models.Address{},
		models.Attribute{},
		models.Brand{},
		models.Cart{},
		models.Category{},
		models.Channel{},
		models.Collect{},
		models.Comment{},
		models.CommentPicture{},
		models.Gallery{},
		models.GoodsAttribute{},
		models.Goods{},
		models.GoodsIssue{},
		models.GoodsRelated{},
		models.GoodsSpecification{},
		models.Keywords{},
		models.Order{},
		models.OrderGoods{},
		models.Product{},
		models.Region{},
		models.SearchHistory{},
		models.Specification{},
		models.Topic{},
		models.UserCoupon{},
		models.User{},
		models.Footprint{},

		//后台表
		models.Admin{},
	)
}
