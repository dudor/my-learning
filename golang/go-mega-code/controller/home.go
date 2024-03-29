package controller

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"gomegacode/model"
	"gomegacode/vm"
	"html/template"
	"log"
	"net/http"
)

type home struct{}

func (h home) registerRouter() {
	//http.HandleFunc("/", middleAuth(indexController))
	//http.HandleFunc("/login", loginController)
	//http.HandleFunc("/logout", middleAuth(logoutController))
	//http.HandleFunc("/register", RegisterController)
	//http.HandleFunc("/user/{username}", middleAuth(ProfileController))

	r := mux.NewRouter()
	r.HandleFunc("/login", loginController)
	r.HandleFunc("/logout", middleAuth(logoutController))
	r.HandleFunc("/register", registerController)
	r.HandleFunc("/user/{username}", middleAuth(profileController))
	r.HandleFunc("/", middleAuth(indexController))
	r.HandleFunc("/profile_edit", middleAuth(profileEditController))
	r.HandleFunc("/follow/{username}", middleAuth(followController))
	r.HandleFunc("/unfollow/{username}", middleAuth(unfollowController))
	r.HandleFunc("/explore", middleAuth(exploreController))
	r.HandleFunc("/reset_password_request", ResetPasswordRequestController)
	r.HandleFunc("/reset_password/{token}", ResetPasswordController)
	http.Handle("/", r)
}

func indexController(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	username, _ := getSessionUser(r)
	page := GetPage(r)
	if r.Method == http.MethodGet {
		fls := GetFlash(w, r)
		log.Println("flash=", fls)
		v := vop.GetVm(username, fls, page, pageLimit)
		templates["index.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		body := r.Form.Get("body")
		err := checkLen("Post", body, 1, 100)
		log.Println("error:=", err)
		if err != "" {
			SetFlash(w, r, err)
		} else {
			err := vm.CreatePost(username, body)
			if err != nil {
				log.Println("Create post error", err)
				w.Write([]byte("Create post error"))
				return
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

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

		errs := CheckLogin(username, password)
		m.AddError(errs...)

		if len(m.Errs) > 0 {
			templates["login.html"].Execute(w, &m)
		} else {
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutController(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func registerController(w http.ResponseWriter, r *http.Request) {
	vop := vm.RegisterViewModelOp{}
	v := vop.GetVm()
	if r.Method == http.MethodGet {
		templates["register.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")
		log.Print("PASSWORD:", pwd1, pwd2)
		errs := CheckRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)
		if len(v.Errs) > 0 {
			templates["register.html"].Execute(w, &v)
		} else {
			if err := AddUser(username, email, pwd1); err != nil {
				log.Print("register user err", err)
				w.Write([]byte("ERROR INSERT DATABASE"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
func profileController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	vop := vm.ProfileViewModelOp{}
	page := GetPage(r)

	v, err := vop.GetVM(sUser, pUser, page, pageLimit)

	fmt.Println(v, err)
	if err != nil {
		msg := fmt.Sprintf("user %s does not exist ", pUser)
		w.Write([]byte(msg))
		return
	}
	templates["profile.html"].Execute(w, &v)
}
func profileEditController(w http.ResponseWriter, r *http.Request) {
	username, _ := getSessionUser(r)
	vop := vm.ProfileEditViewModelOp{}
	v := vop.GetVM(username)
	if r.Method == http.MethodGet {
		templates["profile_edit.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		aboutme := r.Form.Get("aboutme")
		log.Print(aboutme)
		err := model.UpdateAboutMe(username, aboutme)
		if err != nil {
			log.Print("edit profile error:", err)
			w.Write([]byte("Error update aboutme"))
			return
		}
		http.Redirect(w, r, "/user/"+username, 301)
	}
}
func followController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	err := vm.Follow(sUser, pUser)
	if err != nil {
		log.Print("followController Error", err)
		w.Write([]byte("Error in UnFollow"))
		return
	}
	http.Redirect(w, r, "/user/"+pUser, 301)
}
func unfollowController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUser := vars["username"]
	sUser, _ := getSessionUser(r)
	err := vm.UnFollow(sUser, pUser)
	if err != nil {
		log.Print("unfollowController Error", err)
		w.Write([]byte("unfollowController Error"))
		return
	}
	http.Redirect(w, r, "/user/"+pUser, 301)
}
func exploreController(w http.ResponseWriter, r *http.Request) {
	vop := vm.ExploreViewModelOp{}
	username, _ := getSessionUser(r)
	page := GetPage(r)
	v := vop.GetVM(username, page, pageLimit)
	templates["explore.html"].Execute(w, &v)
}
func ResetPasswordRequestController(w http.ResponseWriter, r *http.Request) {
	vop := vm.ResetPasswordRequestViewModelOp{}
	v := vop.GetVmM()
	if r.Method == http.MethodGet {
		templates["reset_password_request.html"].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		email := r.Form.Get("email")
		errs := CheckEmail(email)
		if errs!=""{
			v.AddError(errs)
		}

		if len(v.Errs) > 0 {
			templates["reset_password_request.html"].Execute(w, &v)
		} else {
			log.Println("send email to :", email)
			vop_email := vm.ResetPasswordRequestViewModelOp{}
			v_email := vop_email.GetVmM()
			var contentBytes bytes.Buffer
			tpl, _ := template.ParseFiles("templates/email.html")
			if err := tpl.Execute(&contentBytes, &v_email); err != nil {
				log.Println("Get Parse Template:", err)
				w.Write([]byte("Error send email"))
				return
			}
			content := contentBytes.String()
			go sendEmail(email, "Reset Password", content)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
func ResetPasswordController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token := vars["token"]
	username, err := vm.CheckToken(token)
	if err != nil {
		w.Write([]byte("The token is no longer valid, please go to the login page."))
	}
	vop := vm.ResetPasswordViewModelOp{}
	v := vop.GetVM(token)

	if r.Method == http.MethodGet {
		templates["reset_password.html"].Execute(w, &v)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")
		errs := CheckResetPassword(pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates["reset_password.html"].Execute(w, &v)
		} else {
			if err := vm.ResetUserPassword(username, pwd1); err != nil {
				log.Println("reset User password error:", err)
				w.Write([]byte("Error update user password in database"))
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	}

}
