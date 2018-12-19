package controller

import (
	"gomegacode/vm"
	"log"
	"net/http"
)

type home struct{}

func (h home) registerRouter() {
	http.HandleFunc("/", middleAuth(indexController))
	http.HandleFunc("/login", loginController)
	http.HandleFunc("/logout",middleAuth(logoutController))
	http.HandleFunc("/register",RegisterController)
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

	errs:= CheckLogin(username,password)
	m.AddError(errs...)



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

func RegisterController(w http.ResponseWriter,r *http.Request)  {
	vop := vm.RegisterViewModelOp{}
	v:= vop.GetVm()
	if r.Method == http.MethodGet {
		templates["register.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost{
		r.ParseForm()
		username:= r.Form.Get("username")
		email:=r.Form.Get("email")
		pwd1:= r.Form.Get("pwd1")
		pwd2:= r.Form.Get("pwd2")
		log.Print("PASSWORD:",pwd1,pwd2)
		errs:= CheckRegister(username,email,pwd1,pwd2)
		v.AddError(errs...)
		if len(v.Errs)>0{
			templates["register.html"].Execute(w,&v)
		}else{
			if err:= AddUser(username,email,pwd1);err!=nil{
				log.Print("register user err",err)
				w.Write([]byte("ERROR INSERT DATABASE"))
				return
			}
			setSessionUser(w,r,username)
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}
	}
}