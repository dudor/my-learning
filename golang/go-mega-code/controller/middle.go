package controller

import (
	"fmt"
	"gomegacode/model"
	"log"
	"net/http"
)

func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		if username != "" {
			fmt.Println("update %s last seen time", username)
			model.UpdateLastSeen(username)
		}

		if err != nil {
			log.Print("middle get session ERR , redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			log.Print("middle get session username:", username)
			next.ServeHTTP(w, r)
		}
	}
}
