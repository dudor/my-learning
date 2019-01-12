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
	flashName      string
	pageLimit      int
)

func init() {
	templates = PopulateTemplates()
	store = sessions.NewCookieStore([]byte("secretkeycontents"))
	sessionName = "go-mega"
	flashName = "go-mega-flash"
	pageLimit = 5
}

func Startup() {
	homeController.registerRouter()
}
