package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 添加用户模块路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块路由接口
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块路由接口
		// 添加文章
		auth.POST("article/add", v1.AddArt)
		// 查询所有文章
		// 编辑文章
		auth.PUT("article/:id", v1.EditArt)
		// 删除文章
		auth.DELETE("article/:id", v1.DeleteArt)
	}
	router := r.Group("api/v1")
	{
		// 添加用户
		router.POST("user/add", v1.AddUser)
		// 查询用户列表
		router.GET("users", v1.GetUsers)
		// 查询分类列表
		router.GET("category", v1.GetCate)
		// 查询单个文章信息
		router.GET("article/info", v1.GetArtInfo)
		// 查询分类下所有文章列表
		router.GET("article", v1.GetCateArts)
		//	登录
		router.POST("login", v1.Login)
	}

	_ = r.Run(utils.HttpPost)
}
