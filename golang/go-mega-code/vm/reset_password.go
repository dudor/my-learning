package vm

import "gomegacode/model"

type ResetPasswordViewModelOp struct {
}
type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

func (ResetPasswordViewModelOp) GetVM(token string) ResetPasswordViewModel {
	var v ResetPasswordViewModel
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

func CheckToken(tokenstr string)(string, error){
	return model.CheckToken(tokenstr)
}
func ResetUserPassword(username,password string)error  {
	return model.UpdatePassword(username,password)
}