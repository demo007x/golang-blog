package controllers

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TagIndex(c *gin.Context) {
	var tags []modules.Tag

	err := driver.Db.Find(&tags).Error
	if err != nil {
		utils.RedirectBack(c)
		return
	}
	auth := Auth{}.GetAuth(c)
	data := &struct {
		Auth
		Tags []modules.Tag
	}{
		auth,
		tags,
	}

	c.HTML(http.StatusOK, "tagIndex", data)
}

// 标签
func GetArticleByTagName(c *gin.Context) {
	tagName := c.Param("name")
	if tagName == "" {
		utils.RedirectBack(c)
		return
	}

	// 查找t具有tag的文章
	var articles []modules.Article
	dbQuery := driver.Db.Table("articles").Where("find_in_set('" + tagName + "', tags)").Order("view_num")
	paginate, err := (&modules.Pagination{}).Paginate(c, *dbQuery, &articles)

	if err != nil {
		panic(err)
	}

	auth := Auth{}.GetAuth(c)
	data := struct {
		Paginate modules.Pagination
		Auth
	}{
		*paginate, auth,
	}
	c.HTML(http.StatusOK, "index", data)
}
