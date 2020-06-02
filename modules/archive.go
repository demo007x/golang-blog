package modules

import (
	"github.com/jinzhu/gorm"
)

type Archive struct {
	gorm.Model
	ArchiveDate string `gorm:"char(10)"; not null; default ''`
	ArticleIds string `gorm:"varchar(255); not null; default '';"`
}
