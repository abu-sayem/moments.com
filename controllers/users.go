package controllers

import (
	//"github.com/gorilla/schema"
	"fmt"
	"moments.com/views"
	"net/http"
)
func NewUsers() *Users {
	return &Users{
		 NewView: views.NewView("bootstrap", "users/new"),
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
	Password string `schema: "password"`
}

//POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err !=nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
