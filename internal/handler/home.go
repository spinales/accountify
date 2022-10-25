package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./web/templates/home.page.tmpl",
		"./web/templates/home/navbar.home.tmpl",
		"./web/templates/home/hero.home.tmpl",
		"./web/templates/home/features.home.tmpl",
		"./web/templates/home/pricing.home.tmpl",
		"./web/templates/home/faq.home.tmpl",
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

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./web/templates/contact.page.tmpl",
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

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Dashboard Page!!")
}

func TrasanctionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Trasanctions Page!!")
}
