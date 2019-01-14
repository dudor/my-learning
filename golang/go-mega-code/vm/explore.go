package vm

import "gomegacode/model"

type ExploreViewModelOp struct {

}
type ExploreViewModel struct {
	BaseViewModel
	BasePageViewModel
	Posts []model.Post
}

func (ExploreViewModelOp)GetVM(username string,page,limit int) ExploreViewModel  {
	posts,total,_ := model.GetPostsByPageAndLimit(page,limit)
	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.SetCurrentUser(username)
	v.SetBasePageViewModel(total,page,limit)
	v.Posts = *posts
	return v
}