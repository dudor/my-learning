package vm

import (
	"gomegacode/model"
)

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
	Flash string

	BasePageViewModel
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVm11(username, flash string) IndexViewModel {
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
func (IndexViewModelOp) GetVm(username, flash string, page, limit int) IndexViewModel {
	u, _ := model.GetUserByUsername(username)
	posts,total,_:=u.FollowingPostsByPageAndLimit(page,limit)
	v := IndexViewModel{}
	v.SetTitle("HomePage")
	v.Posts = *posts
	v.Flash = flash
	v.SetCurrentUser(username)
	v.SetBasePageViewModel(total,page,limit)
	return v
}
func CreatePost(username, post string) error {
	u, _ := model.GetUserByUsername(username)
	return u.CreatePost(post)
}
