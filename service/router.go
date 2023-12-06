package service

import (
	"github.com/gin-gonic/gin"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	apiGroup := router.Group("/api/") // 外部调用接口
	v1Group := apiGroup.Group("/v1/") // 版本号

	// 添加中间键
	router.Use(utilityFunction.ReturnHeader()) // 响应头中间件

	v1Group.GET("/article", getFunction.ArticleList)   // 文章列表
	v1Group.GET("/tag", getFunction.TagList)           // 标签列表
	v1Group.GET("/category", getFunction.CategoryList) // 分类列表

	v1Group.GET("/article/:id", getFunction.GetArticle)   // 文章内容
	v1Group.GET("/tag/:id", getFunction.GetTag)           // 标签内容
	v1Group.GET("/category/:id", getFunction.GetCategory) // 标签内容

	v1Group.POST("/tag", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), addFunction.AddTag)           // 添加标签
	v1Group.POST("/category", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), addFunction.AddCategory) // 添加分类
	v1Group.POST("/group", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), addFunction.AddGroup)       // 添加用户组
	v1Group.POST("/user", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), addFunction.AddUser)         // 添加用户
	v1Group.POST("/article", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), addFunction.AddArticle)   // 添加文章

	v1Group.PUT("/tag", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), updateFunction.UpdateTag)           // 更新标签
	v1Group.PUT("/article", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), updateFunction.UpdateArticle)   // 更新文章
	v1Group.PUT("/group", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), updateFunction.UpdateGroup)       // 更新用户组
	v1Group.PUT("/user", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), updateFunction.UpdateUser)         // 更新用户
	v1Group.PUT("/category", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), updateFunction.UpdateCategory) // 更新分类

	v1Group.DELETE("/tag", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), deleteFunction.DeleteTag)           // 删除标签
	v1Group.DELETE("/category", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), deleteFunction.DeleteCategory) // 删除类别
	v1Group.DELETE("/group", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), deleteFunction.DeleteGroup)       // 删除用户组
	v1Group.DELETE("/user", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), deleteFunction.DeleteUser)         // 删除用户
	v1Group.DELETE("/article", utilityFunction.verifyHeaderLoginCode(), utilityFunction.CheckLoginMiddleware(), deleteFunction.DeleteArticle)   // 删除文章

	v1Group.POST("/userLogin", userFunction.UserLogin) // 用户登陆

	// 通用型路由组
	router.GET("/", getFunction.MainPage) // 主页

	router.GET("/about", getFunction.GetAbout) // 关于

	router.POST("/upload", UploadFile) // 上传

	router.GET("/test/:id", test) // 测试

}

func Router(router *gin.Engine) {
	router.NoRoute(utilityFunction.NoFoundPage) // 404错误
	routerMain(router)                          // 主路由
}
