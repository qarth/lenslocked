package controllers

import (
	"fmt"
	"net/http"
	"github.com/qarth/lenslocked/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Email is", form.Email)
	fmt.Fprintln(w, "Password is", form.Password)
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}
