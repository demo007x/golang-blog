package controllers

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
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

// GetTags 获取tags列表
func AjaxTags(c *gin.Context) {
	var tags []modules.Tag

	err := driver.Db.Select("id, name").Find(&tags).Error
	if err != nil {
		response := utils.Response{
			Status: 0,
			Data:   tags,
			Msg:    err.Error(),
		}
		c.JSON(http.StatusOK, response.FailedResponse())
		return
	}

	//response := utils.Response{
	//	Status: 0,
	//	Data:   tags,
	//	Msg:    "",
	//}
	c.JSON(http.StatusOK, tags)
	return
}
