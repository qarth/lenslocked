package controllers

import (
	"fmt"
	"net/http"

	"github.com/qarth/lenslocked/views"
)

func NewGallery() *Galleries {
	return &Galleries{
	}
}

func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>I am a gallery!</h1>")
}

type Galleries struct {
	NewView *views.View
}