package main

import (
	"net/http"

	"github.com/fzzp/ebook/internal/controllers"
)

func adminRouter(repo *controllers.Repository) *http.ServeMux {
	// 管理员路由
	mux := http.NewServeMux()
	mux.HandleFunc("POST /admin/book", repo.PostBook)
	mux.HandleFunc("PUT /admin/book", repo.PutBook)
	mux.HandleFunc("DELETE /admin/book/{id}", repo.DeleteBook)

	return mux
}

func NewRouter(repo *controllers.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	// 非管理员路由
	mux.HandleFunc("GET /home", repo.HomeView)
	mux.HandleFunc("GET /books", repo.ListBooks)

	mux.HandleFunc("GET /signup", repo.SignupView)
	mux.HandleFunc("POST /signup", repo.PostSignUp)

	mux.HandleFunc("GET /login", repo.LoginView)
	mux.HandleFunc("POST /login", repo.PostLogin)

	mux.HandleFunc("PUT /updateMe", repo.PostLogin)

	mux.Handle("/", adminRouter(repo))

	return mux
}
