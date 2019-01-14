package vm

import (
	"gomegacode/model"
)

type EmailViewModelOp struct {

}
type EmailViewModel struct {
	Username string
	Token string
	Server string
}

func (EmailViewModelOp)GetVM(email string)EmailViewModel  {
	var v EmailViewModel
	u,_:= model.GetUserByEmail(email)
	v.Server ="http://localhost:8081"
	v.Token=""
	v.Username = u.Username
	return v
}
