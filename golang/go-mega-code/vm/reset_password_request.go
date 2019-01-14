package vm

import (
	"gomegacode/model"
	"log"
)

type ResetPasswordRequestViewModelOp struct {
}
type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

func (ResetPasswordRequestViewModelOp) GetVmM() ResetPasswordRequestViewModel {
	v:= ResetPasswordRequestViewModel{}
	v.SetTitle("Reset Password Request")
	return v
}
func CheckEmailExist(email string)bool  {
	_,err := model.GetUserByEmail(email)
	if err!=nil{
		log.Print("Cannt find user by email : ",email)
		return false
	}
	return true
}
