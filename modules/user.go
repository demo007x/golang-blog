package modules

import (
	"blog/bootstrap/driver"
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

// 更新用户信息
func (u *User) Update() error  {
	return driver.Db.Save(&u).Error
}

// 用户更新密码
func (u *User) UpdatePassword(password string) (err error) {
	err  = driver.Db.Update(&u).Error
	return
}

// 根据id获取user
func GetUserByID(id int) (User, error)  {
	var user User
	err := driver.Db.First(&user, id).Error
	return user, err
}




