package models
import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectionDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-blog"))
	
	if err!= nil {
		panic(err)
	}

	database.AutoMigrate(&Blog{})

	DB = database
}