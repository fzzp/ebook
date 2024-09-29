package controllers

import (
	"net/http"

	"github.com/fzzp/ebook/views/portal"
)

func (rp Repository) SignupView(w http.ResponseWriter, r *http.Request) {
	portal.SignUp().Render(r.Context(), w)
}

func (rp Repository) PostSignUp(w http.ResponseWriter, r *http.Request) {
}

func (rp Repository) LoginView(w http.ResponseWriter, r *http.Request) {
	portal.Login().Render(r.Context(), w)
}

func (rp Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
}

func (rp Repository) UpdateMe(w http.ResponseWriter, r *http.Request) {
}
