package controllers

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 项目首页
func Index(c *gin.Context) {
	var articles []modules.Article

	dbQuery := driver.Db.Table("articles").Order("created_at desc")
	paginate, err := (&modules.Pagination{}).Paginate(c, *dbQuery, &articles)

	if err != nil {
		panic(err)
	}

	auth := Auth{}.GetAuth(c)
	header := Header{Title: ""}
	data := struct {
		Paginate modules.Pagination
		Auth
		Header
	}{
		*paginate, auth, header,
	}

	c.HTML(http.StatusOK, "index", data)
}
