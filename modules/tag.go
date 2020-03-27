package modules

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(25); not null; default ''; unique"`
	UseNum int `gorm:"type:int(10); not null; default 0"`
}
