package driver

import (
	"blog/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var Db *gorm.DB

func InitConn(str string)  {
	switch str {
	case "mysql":
		Db = mySql()
	default:
		mySql()
	}
}

// mysql connect
func mySql() *gorm.DB {

	database := config.Config.Section("database")
	//dbType, err := database.GetKey("Type")
	//if err != nil {
	//	panic(err)
	//}

	user, err := database.GetKey("User")
	if err != nil {
		panic(err)
	}

	password, err := database.GetKey("Password")
	if err != nil {
		panic(nil)
	}

	host, err := database.GetKey("Host")

	if err != nil {
		panic(err)
	}
	dbName, err := database.GetKey("Name")

	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host,dbName)
	//dsn := "root:YYL521wxl@@tcp(47.105.50.53:3306)/go_blog?charset=utf8&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", "root:YYL521wxl@@tcp(47.105.50.53:3306)/go_blog?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	section,_ := config.Config.GetSection("env")
	mode,_ := section.GetKey("Mode")
	if  mode.String() == "debug" {
		db.LogMode(true)
	}

	return db
}

// Mongodb connect
func mongoDb()  {
	// do some thing
}

// Redis connect
func redis()  {
	// do some thing
}

//Memcache connect
func memcache()  {
	// do some thing
}