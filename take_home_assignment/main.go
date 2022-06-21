package main

import (
	"net/http"

	"github.com/go-playground/log/v8"

	"github.com/tmammado/take-home-assignment/db"
	"github.com/tmammado/take-home-assignment/handler"
	"github.com/tmammado/take-home-assignment/repository"
)

func main() {
	db := db.Init()
	userRepo := repository.NewUserRepo(db)
	app := handler.NewApp(userRepo)
	router := app.NewRouter()

	log.Info("API is running at port 8080")
	log.Error(http.ListenAndServe(":8080", router))
}
