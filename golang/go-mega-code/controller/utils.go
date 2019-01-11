package controller

import (
	"errors"
	"fmt"
	"gomegacode/vm"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic(err)
	}
	fls, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, fl := range fls {
		//fmt.Println(fl.Name())
		f, err := os.Open(basePath + "/content/" + fl.Name())
		if err != nil {
			panic(err)
		}
		log.Print(f.Name())
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic(err)
		}
		result[fl.Name()] = tmpl
	}
	return result
}

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}
	val := session.Values["user"]
	log.Print("var=", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("can not get session user")
	}
	log.Print("username = ", username)
	return username, nil
}
func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func checkLen(fieldName, fieldValue string, minLen, maxLen int) string {
	lenField := len(fieldValue)
	if lenField < minLen {
		return fmt.Sprintf("%s field is too short, less than %d", fieldName, minLen)
	}
	if lenField > maxLen {
		return fmt.Sprintf("%s field is too long, more than %d", fieldName, maxLen)
	}
	return ""
}
func CheckUsername(username string) string {
	return checkLen("username", username, 6, 20)
}
func CheckPassword(password string) string {
	return checkLen("password", password, 6, 20)
}
func CheckEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return fmt.Sprintf("Email field not a valid email")
	}
	return ""
}
func CheckUserPassword(username, password string) string {
	if vm.CheckLogin(username, password) == false {
		return fmt.Sprintf("Username and Password is not correct")
	}
	return ""
}
func CheckUserExist(username string) string {
	if vm.CheckUserExist(username) == false {
		return fmt.Sprintf("username already exsit,please choose another ")
	}
	return ""
}
func CheckLogin(username, password string) []string {
	var errs []string
	if err := CheckUsername(username); len(err) > 0 {
		errs = append(errs, err)
	}
	if err := CheckPassword(password); len(err) > 0 {
		errs = append(errs, err)
	}
	if err := CheckUserPassword(username, password); len(err) > 0 {
		errs = append(errs, err)
	}
	return errs
}
func CheckRegister(username, email, password1, password2 string) []string {
	var errs []string
	if password1 != password2 {
		errs = append(errs, "password does not match")
	}
	if err := CheckUsername(username); len(err) > 0 {
		errs = append(errs, err)
	}
	if err := CheckPassword(password1); len(err) > 0 {
		errs = append(errs, err)
	}
	if err := CheckEmail(email); len(err) > 0 {
		errs = append(errs, err)
	}
	if err := CheckUserExist(username); len(err) > 0 {
		errs = append(errs, err)
	}
	return errs
}
func AddUser(username, email, password string) error {
	return vm.AddUser(username, password, email)
}
func SetFlash(w http.ResponseWriter,r *http.Request,content string)  {
	session,_:= store.Get(r,sessionName)
	session.AddFlash(content,flashName)
	session.Save(r,w)
}
func GetFlash(w http.ResponseWriter,r *http.Request) string {
	session,_:= store.Get(r,sessionName)
	content:= session.Flashes(flashName)
	if content == nil{
		return ""
	}
	session.Save(r,w)
	return fmt.Sprintf("%v",content)
}