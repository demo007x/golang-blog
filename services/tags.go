package services

import (
	"blog/bootstrap/driver"
	"blog/modules"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

// handleTags 处理文章的tags
func HandleTags(tag string) bool  {
	var tagSlice []string
	tagSlice = strings.Split(tag, ",")
	// 查找表里面所有的tags, 如果没有找到则添加， 如果找到了useNum + 1
	var tags[]modules.Tag
	err := driver.Db.Where("name in (?)", tagSlice).Find(&tags).Error
	if err != nil {
		log.Println(err)
		return false
	}

	// 将 tagSlice 转换为map
	tagMap := make(map[string]string)
	for _, ts := range tagSlice {
		tagMap[ts] = ts
	}

	var tagsStructSlice []modules.Tag

	for _, tm := range tagMap {
		var ts = modules.Tag{
			Name:   tm,
			UseNum: 1,
		}
		for _, tag := range tags {
			if tm == tag.Name {
				// 找到的情况
				ts = modules.Tag{
					Model:  gorm.Model{
						ID:tag.ID,
					},
					Name:   tag.Name,
					UseNum: tag.UseNum + 1,
				}
			}
		}
		tagsStructSlice = append(tagsStructSlice, ts)
	}

	// 处理数据 获取插入数据库
	for _, v :=  range tagsStructSlice {
		var err error
		if v.ID == 0 {
			err = driver.Db.Create(&v).Error
			if err != nil {
				panic(err)
			}
		} else {
			err = driver.Db.Model(modules.Tag{}).Update(v).Error
		}
		if err != nil {
			return false
		}
	}
	return  true
}
