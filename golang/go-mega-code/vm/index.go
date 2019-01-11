package vm

import (
	"gomegacode/model"
)

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVm(username,flash string) IndexViewModel {
	u1, err := model.GetUserByUsername(username)
	if err != nil {
		return IndexViewModel{}
	}
	posts, _ := u1.FollowingPosts()
	v := IndexViewModel{}
	v.Title = "Index Page"
	v.Posts = *posts
	v.SetCurrentUser(username)
	v.Flash = flash
	return v
}
func CreatePost(username,post string) error  {
	u,_:= model.GetUserByUsername(username)
	return u.CreatePost(post)
}