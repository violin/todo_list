package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.LogMode(true)
	if gin.Mode() == gin.ReleaseMode {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	fmt.Println("complete mysql initialize")

	migration()
	fmt.Println("complete mysql migration")

}
