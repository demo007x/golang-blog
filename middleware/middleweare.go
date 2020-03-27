package middleware

import (
	"blog/controllers"
	"blog/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// c中注入auth信心
func SetAuth(c *gin.Context)  {
	sess := sessions.Default(c)
	auth := sess.Get("auth")

	if auth != nil {
		c.Set("auth", auth)
	}

	c.Next()
}

// Authorization 验证用户
func Authorization(c *gin.Context)  {
	auth := (&controllers.Auth{}).GetAuth(c)
	if auth.Id == 0 {
		// 用户未登录情况
		utils.Redirect(c, "/login")
		return
	}
	c.Next()
}
