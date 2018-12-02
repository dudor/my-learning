package main

import (
	"gomegacode/model"
	"net/http"
	"gomegacode/controller"
)

func main() {

	controller.Startup()
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}



