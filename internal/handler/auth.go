package handler

import (
	"html/template"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/login.page.tmpl",
		"./web/templates/base.layout.tmpl",
		"./web/templates/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/signup.page.tmpl",
		"./web/templates/base.layout.tmpl",
		"./web/templates/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func RecoveryHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/recovery.page.tmpl",
		"./web/templates/base.layout.tmpl",
		"./web/templates/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
