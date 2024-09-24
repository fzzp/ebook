package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fzzp/ebook/components"
)

func main() {
	fmt.Println("Hello Go!")

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if err := components.AdminLayout().Render(r.Context(), w); err != nil {
			slog.Error("组件渲染失败", slog.String("error", err.Error()))
		}
	})

	mux.HandleFunc("/showimg", func(w http.ResponseWriter, r *http.Request) {
		components.ShowImg().Render(r.Context(), w)
	})

	http.ListenAndServe(":9400", mux)
}
