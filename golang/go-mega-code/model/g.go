package model

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func SetDB(database *gorm.DB)  {
	db = database
}
func ConnectToDB() *gorm.DB  {
	connectionStr := "../gomega.db"
	log.Print("connect to db ...")
	db,err:= gorm.Open("sqlite3",connectionStr)
	if err!=nil{
		panic(err)
	}
	db.SingularTable(true)
	return db
}