package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Username string
}
type Post struct {
	User
	Body string
}

type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		posts := []Post{
			Post{Body: "body1", User: User{Username: "user1"}},
			Post{Body: "body2", User: User{Username: "user2"}},
			Post{Body: "body3", User: User{Username: "user3"}},
		}

		indexVM := IndexViewModel{
			Title: "Homepage",
			User:  User{Username: "dudor"},
			Posts: posts,
		}

		tpls := PopulateTemplates()
		tpls["index.html"].Execute(w, &indexVM)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

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
		fmt.Println(fl.Name())
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
