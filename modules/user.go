package modules

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(25)"`
	Password string `gorm:"type:varchar(32); not null; default ''"`
	Salt string `gorm:"type:char(4); size:4; not null; default ''"`
	Email string `gorm:"type:varchar(100);unique_index"`
	Profession string `gorm:"type:varchar(255); not null; default ''"`
	Avatar string `gorm:"type:varchar(255); not null; default ''"`
	Articles []Article // 用户有多篇文章
}




