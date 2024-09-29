package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dbrepo "github.com/fzzp/ebook/dbrepo/sqlc"
	"github.com/fzzp/ebook/internal/controllers"
	"github.com/fzzp/ebook/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// 加载配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 设置 json 格式 slog
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 连接 mysql
	conn, err := openDB()
	if err != nil {
		slog.Error("打开数据库失败: " + err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxIdleTime(25 * time.Minute)

	// 创建service层
	store := dbrepo.NewSQLStore(conn)
	userSvc := services.NewUserService(store)

	repo := controllers.NewRepository(userSvc)

	mux := NewRouter(repo)

	fileServer := http.FileServer(http.Dir("./views/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("WEB_PORT")),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  2 * time.Minute,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	go func() {
		slog.Info("start on " + server.Addr)
		err = server.ListenAndServe()
		if err != nil {
			slog.Error("server.ListenAndServe error: " + err.Error())
			os.Exit(1)
		}
	}()

	// 优雅关机
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-quitChan

	slog.Info("获取到系统退出信号", slog.Any("sig", sig))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("server.Shutdown error: " + err.Error())
		os.Exit(1)
	}

}

func openDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		"ebook_mysql", // 容器名
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
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
