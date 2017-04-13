package controllers

import "github.com/qarth/lenslocked/views"

func NewStatic() *Static {
	return &Static {
		Home: views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		Faq: views.NewView("boostrap", "static/faq"),
	}
}

type Static struct {
	Home *views.View
	Contact *views.View
	Faq *views.View
}