package vm

import (
	"gomegacode/model"
	"log"
)

type RegisterViewModel struct {
	LoginViewModel
}
type RegisterViewModelOp struct {
}

func (RegisterViewModelOp) GetVm() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

func CheckUserExist(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Print("Cannot find username :", username)
		return true
	}
	return false
}

func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}
