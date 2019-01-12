package main

import (
	"github.com/gorilla/context"
	"gomegacode/controller"
	"gomegacode/model"
	"log"
	"net/http"
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
	return
	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("username1", "password1", "email1@qq.com")
	model.AddUser("username2", "password2", "email2@qq.com")

	u1, _ := model.GetUserByUsername("username1")
	u1.CreatePost("this is my first post content")
	model.UpdateAboutMe("username1", "i am username1 ,this is about me content")

	u2, _ := model.GetUserByUsername("username2")
	u2.CreatePost("today is a good day")
	model.UpdateAboutMe("username2", "i am username2, please follow me ")

	u2.Follow("username1")

	/*
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
	*/
}
