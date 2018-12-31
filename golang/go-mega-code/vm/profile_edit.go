package vm

import "gomegacode/model"

type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser model.User
}

type ProfileEditViewModelOp struct {
}

func (ProfileEditViewModelOp) GetVM(username string) ProfileEditViewModel {
	v := ProfileEditViewModel{}
	u, _ := model.GetUserByUsername(username)
	v.ProfileUser = *u
	v.SetTitle("Edit Profile Page")
	v.SetCurrentUser(username)
	return v
}
func UpdateAboutMe(username, aboutme string) error {
	return model.UpdateAboutMe(username, aboutme)
}
