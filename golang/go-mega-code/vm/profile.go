package vm

import "gomegacode/model"

type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post
	ProfileUser    model.User
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	BasePageViewModel
}

type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetVM(curUser, pUser string,page,limit int) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("ProfilePage")
	u1, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts,total, _ := model.GetPostsByUserIDPageAndLimit(u1.ID,page,limit)
	v.Posts = *posts
	v.ProfileUser = *u1
	v.SetCurrentUser(curUser)
	v.Editable = (curUser == pUser)
	v.SetBasePageViewModel(total,page,limit)
	if !v.Editable {
		v.IsFollow = u1.IsFollowedByUser(curUser)
	}
	v.FollowersCount = u1.FollowersCount()
	v.FollowingCount = u1.FollowingCount()
	return v, nil
}
func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Follow(b)
}
func UnFollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.UnFollow(b)
}
