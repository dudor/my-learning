package main

import (
	"./controller"
	"net/http"
)

func main() {
	controller.Startup()
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}



