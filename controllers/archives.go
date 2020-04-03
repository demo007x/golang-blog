package controllers

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 归档
func Archives(c *gin.Context) {
	auth := Auth{}.GetAuth(c)
	var archives []modules.Archive
	err := driver.Db.Find(&archives).Error
	if err != nil {
		log.Fatalln(err)
	}

	type articleItems []modules.Article
	Archives := make(map[string]articleItems)
	if len(archives) > 0 {
		for _, v := range archives {
			// 查找文章
			var ids []int
			for _, id := range strings.Split(v.ArticleIds, ",") {
				id, _ := strconv.Atoi(id)
				ids = append(ids, id)
			}
			var Items articleItems
			driver.Db.Table("articles").Where("id in (?)", ids).Find(&Items)
			Archives[v.ArchiveDate] = Items
		}
	}
	for _, v := range Archives {
		for _, item := range v {
			fmt.Printf("%#v", item)
		}
	}
	header := Header{Title: "文章归档"}
	data := struct {
		Auth
		Archives map[string]articleItems
		Header   Header
	}{auth, Archives, header}
	c.HTML(http.StatusOK, "archives", data)
}
