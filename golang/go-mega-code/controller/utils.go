package controller

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
