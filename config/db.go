package config
import (
	"fmt"
	"os"
_"github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)


func connection() *gorm.DB{
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	URL:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",USER,PASS,HOST,PORT,DB_NAME)

	db, err := gorm.Open("mysql", URL)
	
	if err!= nil {
		panic(err.Error())
	}

	return db
}