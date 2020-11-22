package controllers

import (
	"github.com/gorilla/schema"
	"fmt"
	"moments.com/views"
	"net/http"
)
func NewUsers() *Users {
	return &Users{
		 NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// GRT /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type SignupForm struct {
	Email string `schema: "email"`
	Password string `svhema: "password"`
}

//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err !=nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
}
