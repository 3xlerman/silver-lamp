package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:lermanworks@gmail.com\">lermanworks@gmail.com")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Ask me something if you want")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq",faq)
	http.ListenAndServe(":3000", r)
}
