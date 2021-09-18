package routers

import (
	"github.com/gin-gonic/gin"
	adminAd "shop_mater/admin/ad"
	adminAdPosition "shop_mater/admin/ad.position"
	adminCategory "shop_mater/admin/category"
	adminGoods "shop_mater/admin/goods"
	adminLogin "shop_mater/admin/login"
	"shop_mater/admin/upload"
	adminUser "shop_mater/admin/user"
	"shop_mater/service/login"
)

// AdminRouters 后台路由
func AdminRouters(r *gin.Engine) {
	//路由分组
	v2 := r.Group("/admin")
	//后台中间件 token校验
	v2.Use(login.AdminMiddleware)

	//用户登录
	v2.POST("/user/login", adminLogin.Login)

	//产品模块
	v2.POST("/goods/list", adminGoods.List)
	v2.POST("/goods/delete", adminGoods.Delete)
	v2.POST("/goods/create", adminGoods.Create)
	v2.POST("/goods/update", adminGoods.Update)

	//分类模块
	v2.POST("/category/list", adminCategory.List)
	v2.POST("/category/delete", adminCategory.Delete)
	v2.POST("/category/create", adminCategory.Create)
	v2.POST("/category/update", adminCategory.Update)

	//广告模块
	v2.POST("/ad/list", adminAd.List)
	v2.POST("/ad/delete", adminAd.Delete)
	v2.POST("/ad/create", adminAd.Create)
	v2.POST("/ad/update", adminAd.Update)
	v2.POST("/adposition/list", adminAdPosition.List)

	//用户模块
	v2.POST("/user/list", adminUser.List)

	//图片上传
	v2.POST("/picture/upload", upload.Upload)
}
