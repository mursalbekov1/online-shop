package app

import (
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"online-shop/internal/config"
	"online-shop/internal/db"
	"online-shop/internal/repository/postgres"
	"online-shop/internal/route"
	"os"
)

func Run() {
	cfg := config.MustLoad()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	logger.Info(cfg.HttpServer.Port)
	logger.Info(cfg.HttpServer.Host)

	db := db.Connect()

	h := postgres.NewUserRepository(db)

	r := route.Router(h)
	err := http.ListenAndServe(cfg.HttpServer.Host+":"+cfg.HttpServer.Port, r)
	if err != nil {
		panic(err)
	}
}
