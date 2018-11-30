package vm

import "../model"

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVm() IndexViewModel {
	u1 := model.User{Username: "user1"}
	u2 := model.User{Username: "user2"}

	posts := []model.Post{
		model.Post{User: u1, Body: "content1"},
		model.Post{User: u2, Body: "content2"},
	}
	return IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
}
