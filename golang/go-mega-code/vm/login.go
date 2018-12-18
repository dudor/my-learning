package vm

import (
	"gomegacode/model"
	"log"
)

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct {
}

func (LoginViewModelOp) GetVm() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Print("can not find username:",username)
		log.Print(err)
		return false
	}
	return user.CheckPassword(password)
}
