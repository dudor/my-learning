package vm

import (
	"gomegacode/model"
)

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVm(username string) IndexViewModel {
	u1, err := model.GetUserByUsername(username)
	if err != nil {
		return IndexViewModel{}
	}
	posts, _ := model.GetPostByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	v.SetCurrentUser(username)
	return v
}
