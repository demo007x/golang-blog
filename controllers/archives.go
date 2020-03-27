package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 归档
func Archives(c *gin.Context) {
	c.HTML(http.StatusOK, "archives", nil)
}
