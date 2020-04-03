package routers

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func Home(r *gin.Engine) {
	home := r.Group("/")
	{
		// 首页
		home.GET("/", controllers.Index)

		article := home.Group("/article", middleware.Authorization)
		{
			// 创建文章
			article.GET("/user", controllers.UserArticleList)
			article.GET("/create", controllers.CreateArticle)
			article.POST("/create", controllers.SaveArticle)
			article.GET("/edit/:id", controllers.EditArticle)
			article.GET("/delete/:id", controllers.DelArticle)
		}

		// 文章详情
		home.GET("/detail/:id", controllers.Detail)

		// 标签页面
		tag := home.Group("/tags")
		{
			tag.GET("/", controllers.TagIndex)
			tag.GET("/title/:name", controllers.GetArticleByTagName)
			tag.GET("/ajax/list", controllers.AjaxTags)
		}

		home.GET("/archives", controllers.Archives)
		// 注册
		home.GET("/join", controllers.Register)
		home.POST("/join", controllers.DoRegister)

		// sign in
		home.GET("/login", controllers.Login)
		home.POST("/login", controllers.DoLogin)
		home.GET("/logout", controllers.Logout)
	}
}
