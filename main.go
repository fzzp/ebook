package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/fzzp/ebook/components"
	dbrepo "github.com/fzzp/ebook/dbrepo/sqlc"
	"github.com/fzzp/ebook/formdto"
	"github.com/fzzp/ebook/services"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

var (
	formDec = form.NewDecoder()
)

func main() {
	fmt.Println("Hello Go!")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"ebook_scret",
		"ebook_mysqldb_1",
		3306,
		"ebook",
	)

	conn, err := openDB(dsn)
	if err != nil {
		panic(err)
	}

	store := dbrepo.NewSQLStore(conn)
	svc := services.NewUserService(store)

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

	mux.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}

		var input formdto.CreateUserRequest
		if err = formDec.Decode(&input, r.PostForm); err != nil {
			panic(err)
		}

		err = svc.CreateUser(r.Context(), input)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
			return
		}

		fmt.Fprintf(w, "%s", "注册成功")
	})

	fmt.Println("start on 9400")
	http.ListenAndServe(":9400", mux)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
