package services

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// insertArchive 将发布的文章归档
func SetArticleArchive(article *modules.Article) {
	var archive modules.Archive
	layout := "2006-01-02 03:04:05"
	archiveDateParse, err := time.Parse(layout, article.CreatedAt.Format(layout))
	if err != nil {
		return
	}
	// 查找是否有当前的归档
	theArticleArchiveDate := archiveDateParse.Format("2006-01")
	driver.Db.Where("archive_date = ?", theArticleArchiveDate).First(&archive)
	if archive.ID == 0 {
		// 创建
		archive.ArchiveDate = archiveDateParse.Format("2006-01")
		archive.ArticleIds = fmt.Sprintf("%d", article.ID)
		err := driver.Db.Create(&archive).Error
		if err != nil {
			log.Println(err)
		}
		return
	}

	// update
	ids := archive.ArticleIds
	idsSlice := strings.Split(ids, ",")
	hasTheID := false
	for _, id := range idsSlice {
		nId, _ := strconv.Atoi(id)
		if uint(nId) == article.ID {
			hasTheID = true
			return
		}
	}

	if !hasTheID {
		idsSlice = append(idsSlice, strconv.Itoa(int(article.ID)))
	}
	archive.ArticleIds = strings.Join(idsSlice, ",")
	err = driver.Db.Save(&archive).Error
	if err != nil {
		log.Println(err)
	}
}
