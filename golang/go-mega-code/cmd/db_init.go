package main

import (
	"gomegacode/model"
)

func main() {

	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	db.DropTableIfExists(&model.User{}, &model.Post{})
	db.CreateTable(&model.User{}, &model.Post{})

}
