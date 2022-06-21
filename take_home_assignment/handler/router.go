package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tmammado/take-home-assignment/middleware"
)

func (app *App) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signup", app.Signup).Methods(http.MethodPost)
	router.HandleFunc("/login", app.Login).Methods(http.MethodPost)
	router.HandleFunc("/users", middleware.JwtAuthentication(http.HandlerFunc(app.GetAllUsers)).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/users", middleware.JwtAuthentication(http.HandlerFunc(app.UpdateUser)).ServeHTTP).Methods(http.MethodPut)

	return router
}
