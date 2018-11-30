package controller

import (
	"html/template"
	"io/ioutil"
	"os"
	"fmt"
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
