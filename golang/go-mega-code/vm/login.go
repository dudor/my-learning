package vm

type LoginViewModel struct {
	BaseViewModel
}

type LoginViewModelOp struct {
	
}

func (LoginViewModelOp)GetVm()LoginViewModel  {
	v:= LoginViewModel{}
	v.SetTitle("Login")
	return v
}