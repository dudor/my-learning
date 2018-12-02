package vm

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
