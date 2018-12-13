package controller

import (
	"gomegacode/vm"
	"net/http"
)

type home struct{}

func (h home) registerRouter() {
	http.HandleFunc("/", indexController)
	http.HandleFunc("/login", loginController)
}

func indexController(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	m := vop.GetVm()
	templates["index.html"].Execute(w, &m)
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
		if check(username, password) == false {
			m.AddError("username password not correct,please input again")
		}

		if len(m.Errs) > 0 {
			templates["login.html"].Execute(w, &m)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
func check(username, password string) bool {
	if username == "user1" && password == "password1" {
		return true
	}
	return false
}
