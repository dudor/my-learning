package vm

import "gomegacode/model"

type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	ProfileUser model.User
	Editable bool
}

type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetVM(curUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("ProfilePage")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostByUserID(u1.ID)
	v.Posts = *posts
	v.ProfileUser = *u1
	v.SetCurrentUser(curUser)
	v.Editable = (curUser == pUser)
	return v, nil
}
