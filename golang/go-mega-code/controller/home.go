package controller

import (
	"gomegacode/vm"
	"net/http"
)

type home struct{}

func (h home) registerRouter() {
	http.HandleFunc("/", middleAuth(indexController))
	http.HandleFunc("/login", loginController)
	http.HandleFunc("/logout",middleAuth(logoutController))
}

func indexController(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	username,_:= getSessionUser(r)
	v:= vop.GetVm(username)
	templates["index.html"].Execute(w, &v)


}

func loginController(w http.ResponseWriter, r *http.Request) {
	vop := vm.LoginViewModelOp{}
	m := vop.GetVm()
	if r.Method == http.MethodGet {
		templates["login.html"].Execute(w, &m)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		//fmt.Fprintf(w, "username=%s password=%s", username, password)

		if len(username) < 3 {
			m.AddError("username length must longer than 3")
		}
		if len(password) < 6 {
			m.AddError("password length must longer than 6")
		}

		if vm.CheckLogin(username,password) == false{
			m.AddError("username password not correct,please input again")
		}

		if len(m.Errs) > 0 {
			templates["login.html"].Execute(w, &m)
		} else {
			setSessionUser(w,r,username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutController(w http.ResponseWriter,r *http.Request)  {
	clearSession(w,r)
	http.Redirect(w,r, "/login",http.StatusTemporaryRedirect)
}
