package controllers

import (
	"github.com/gin-gonic/gin"
)

// session 中 auth
type Auth struct {
	Id         int
	Name       string
	Avatar     string
	Profession string
}

// 获取用户登录信息 auth 的方法
func (a Auth) GetAuth(c *gin.Context) Auth {
	auth, exists := c.Get("auth")
	if !exists {
		auth = Auth{
			Id:         0,
			Name:       "",
			Avatar:     "",
			Profession: "",
		}
	}
	return auth.(Auth)
}

type Header struct {
	Title string
}
