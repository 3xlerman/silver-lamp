package main

import (
	"fmt"
	"html/template"
	"net/http"

	"silver-lamp/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var homeTemplate, contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Ask me something if you want")
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Sorry, the page you looking for is not found :(")
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	var h http.Handler = http.HandlerFunc(notfound)
	r.HandleFunc("notfound", notfound)
	r.NotFoundHandler = h

	http.ListenAndServe(":3000", r)
}
