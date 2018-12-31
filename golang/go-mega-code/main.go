package main

import (
	"fmt"
	"gomegacode/controller"
	"gomegacode/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

func main() {
	initDB()
	controller.Startup()
	log.Print("server runing at 127.0.0.1:8081")
	err := http.ListenAndServe(":8081", context.ClearHandler(http.DefaultServeMux))

	if err != nil {
		panic(err)
	}

}

func initDB() {

	db := model.ConnectToDB()
	//defer db.Close()
	model.SetDB(db)
	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "username1",
			PasswordHash: model.GeneratePasswordHash("password1"),
			Email:"email1@qq.com",
			Avatar:fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("email1@qq.com")),
			Posts: []model.Post{
				{
					Body:      "this is post1 contents",
					Timestamp: time.Now(),
				},
				{
					Body:      "this is post2 contents",
					Timestamp: time.Now(),
				},
			},
		},
		{
			Username:     "username2",
			PasswordHash: model.GeneratePasswordHash("password2"),
			Email:"email1@qq.com",
			Avatar:fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("email1@qq.com")),
			Posts: []model.Post{
				{
					Body:      "this is post2 contents",
					Timestamp: time.Now(),
				},
				{
					Body:      "this is post2 contents",
					Timestamp: time.Now(),
				},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
