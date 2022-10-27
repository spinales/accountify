package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spinales/accountify/internal/handler"
)

func main() {
	r := mux.NewRouter()
	fileserver := http.FileServer(http.Dir("./web/css/"))

	// handlers routes
	r.HandleFunc("/", handler.HomeHandler)
	r.HandleFunc("/login", handler.LoginHandler)
	r.HandleFunc("/signup", handler.SignUpHandler)
	r.HandleFunc("/contact", handler.ContactHandler)
	r.HandleFunc("/recovery", handler.RecoveryHandler)
	// r.Handle("/static/", http.StripPrefix("/static", fileserver))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileserver))
	// r.Handle("/static/", http.StripPrefix("/static", fileserver))

	// start server
	log.Println("App running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println(err)
	}
}
