package routers

import (
	"github.com/gin-gonic/gin"
	"shop_mater/controllers/address"
	"shop_mater/controllers/auth"
	"shop_mater/controllers/brand"
	"shop_mater/controllers/cart"
	"shop_mater/controllers/catalog"
	"shop_mater/controllers/collect"
	"shop_mater/controllers/comment"
	"shop_mater/controllers/footprint"
	"shop_mater/controllers/goods"
	"shop_mater/controllers/index"
	"shop_mater/controllers/order"
	"shop_mater/controllers/pay"
	"shop_mater/controllers/region"
	"shop_mater/controllers/search"
	"shop_mater/controllers/topic"
	"shop_mater/service/login"
)

func InitRouters(r *gin.Engine) {
	//路由分组
	v1 := r.Group("/api")
	v1.Use(login.Middleware)

	//首页数据
	v1.GET("/index/index", index.IndexIndex)

	//微信登录
	v1.POST("/auth/loginByWeixin", auth.LoginByWeixin)

	//商品接口
	v1.GET("/goods/count", goods.Count)
	v1.GET("/goods/category", goods.Category)
	v1.GET("/goods/list", goods.List)
	v1.GET("/goods/detail", goods.Detail)
	v1.GET("/goods/hot", goods.Hot)
	v1.GET("/goods/new", goods.New)
	v1.GET("/goods/related", goods.Related)

	//分类数据
	v1.GET("/catalog/index", catalog.Index)
	v1.GET("/catalog/current", catalog.Current)

	//品牌
	v1.GET("/brand/list", brand.List)
	v1.GET("/brand/detail", brand.Detail)

	//专题
	v1.GET("/topic/list", topic.List)
	v1.GET("/topic/detail", topic.Detail)
	v1.GET("/topic/related", topic.Related)

	//购物车数据
	v1.GET("/cart/index", cart.Index)
	v1.GET("/cart/checkout", cart.CheckOut)
	v1.GET("/cart/goodscount", cart.Count)
	v1.POST("/cart/update", cart.Update)
	v1.POST("/cart/checked", cart.Checked)
	v1.POST("/cart/delete", cart.Delete)
	v1.POST("/cart/add", cart.Add)

	//搜索
	v1.GET("/search/index", search.Index)
	v1.GET("/search/helper", search.Helper)
	v1.POST("/search/clearhistory", search.Clear)

	//评论
	v1.GET("/comment/list", comment.List)
	v1.GET("/comment/count", comment.Count)
	v1.POST("/comment/post", comment.Post)

	//收藏
	v1.POST("/collect/addordelete", collect.AddOrDelete)
	v1.GET("/collect/list", collect.List)

	//获取区域列表
	v1.GET("/region/list", region.List)

	//足迹
	v1.GET("footprint/list", footprint.List)
	v1.POST("/footprint/delete", footprint.Delete)

	//收货地址
	v1.GET("/address/list", address.List)
	v1.GET("/address/detail", address.Detail)
	v1.POST("/address/save", address.Save)
	v1.POST("/address/delete", address.Delete)

	//订单
	v1.POST("/order/submit", order.Submit)
	v1.POST("/order/cancel", order.Cancel)
	v1.GET("/order/list", order.List)
	v1.GET("/order/detail", order.Detail)

	//支付
	v1.GET("/pay/prepay", pay.Prepay)
}
