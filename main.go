package main

import (
	"fmt"
	"net/http"

	"github.com/qarth/lenslocked/controllers"

	"github.com/gorilla/mux"
)

func main() {
	galleryC := controllers.NewGallery()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/users/new", galleryC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(errorpage)
	http.ListenAndServe(":3000", r)
}

func errorpage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404</h1>")
}