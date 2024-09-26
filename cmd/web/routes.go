package main

import (
	"fmt"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// 非管理员路由
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.URL.Path)
	})

	// 管理员路由
	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST book", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.URL.Path)
	})

	mux.Handle("/", adminRouter)

	return mux
}
