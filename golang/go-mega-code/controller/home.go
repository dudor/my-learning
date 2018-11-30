package controller

import (
	"fmt"
	"net/http"
	"../vm"
)

type home struct{}

func (h home) registerRouter() {
	http.HandleFunc("/", indexController)
	http.HandleFunc("/login",loginController)
}

func indexController(w http.ResponseWriter, r *http.Request) {
	vop:= vm.IndexViewModelOp{}
	m:= vop.GetVm()
	templates["index.html"].Execute(w,&m)
}

func loginController(w http.ResponseWriter, r *http.Request)  {
	vop := vm.LoginViewModelOp{}
	m:= vop.GetVm()
	if r.Method == http.MethodGet{
		templates["login.html"].Execute(w,&m)
	}
	if r.Method == http.MethodPost{
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		fmt.Fprintf(w,"username=%s password=%s",username,password)
	}

}