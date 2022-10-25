package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spinales/accountify/internal/handler"
)

func main() {
	r := mux.NewRouter()

	// handlers routes
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/login", handler.LoginHandler)
	r.HandleFunc("/signup", handler.SignUpHandler)
	r.HandleFunc("/contact", handler.ContactHandler)
	r.HandleFunc("/recovery", handler.RecoveryHandler)

	// start server
	log.Println("App running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
