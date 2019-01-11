package controller

import (
	"html/template"

	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store          *sessions.CookieStore
	flashName string
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("secretkeycontents"))
	sessionName = "go-mega"
	flashName = "go-mega-flash"
}

func Startup() {
	homeController.registerRouter()
}
