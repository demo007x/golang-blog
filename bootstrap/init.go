package bootstrap

import (
	"blog/controllers"
	"blog/middleware"
	"blog/routers"
	"blog/utils"
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//var App *gin.Engine

// 定义session的健
const COOKIE_SESSION_KEY = "blog_session"

// 初始化项目
func Init() *gin.Engine {
	App := gin.Default()

	// cookie and session
	gob.Register(controllers.Auth{})
	store := cookie.NewStore([]byte("secret"))
	App.Use(sessions.Sessions(COOKIE_SESSION_KEY, store))

	// 模板中添加函数
	App.SetFuncMap(template.FuncMap{
		"html":          utils.Html,
		"tagString2Map": utils.TagString2Map,
		"setLinkTitle": utils.SetLinkTitle,
		"appUrl": utils.AppUrl,
	})
	// 设置模板解析路径
	App.LoadHTMLGlob("./views/**/*")
	// 设置静态文件
	App.Static("/static", "./static")

	// 设置用户信息
	App.Use(middleware.SetAuth)
	App.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404", nil)
	})
	// 注册路由
	routers.Api(App)
	routers.Home(App)

	return App
}
