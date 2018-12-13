package vm

import (
	"gomegacode/model"
	"log"
)

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVm() IndexViewModel {
	/*
		u1 := model.User{Username: "user1"}
		u2 := model.User{Username: "user2"}

		posts := []model.Post{
			model.Post{User: u1, Body: "content1"},
			model.Post{User: u2, Body: "content2"},
		}
		u3, err := model.GetUserByUsername("user1")
		if err != nil {
			log.Print("err:=", err)
			return IndexViewModel{}
		}
		log.Print(u1, posts, u3)
		return IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	*/

	u1, _ := model.GetUserByUsername("user1")
	posts, _ := model.GetPostByUserID(u1.ID)
	log.Println(u1)
	log.Println(posts)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	log.Print(v)
	return v

}
