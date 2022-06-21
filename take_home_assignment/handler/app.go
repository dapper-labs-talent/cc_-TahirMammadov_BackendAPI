package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/log/v8"

	"github.com/tmammado/take-home-assignment/jwt"
	"github.com/tmammado/take-home-assignment/model"
	"github.com/tmammado/take-home-assignment/shared"
)

type IUserDB interface {
	GetUserByEmail(context.Context, string) (*model.User, error)
	CreateUser(context.Context, string, string, string, string) (*model.User, error)
	GetAllUsers(context.Context) (*[]model.User, error)
	UpdateUserInfo(context.Context, string, string, string) (*model.User, error)
}

type App struct {
	repo IUserDB
}

func NewApp(repo IUserDB) *App {
	return &App{
		repo: repo,
	}
}

func (app *App) Signup(res http.ResponseWriter, req *http.Request) {

	var createUserRequest shared.CreateUserRequest

	// Decode Request
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&createUserRequest); err != nil {
		log.Error("Error while decoding request body: ", err.Error())
		app.respondWithError(res, http.StatusBadRequest, "Cannot decode request body")
		return
	}

	// Validate Request
	err := createUserRequest.Validate()
	if err != nil {
		app.respondWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	// Check if user exists
	userModel, _ := app.repo.GetUserByEmail(req.Context(), createUserRequest.Email)
	if userModel != nil {
		app.respondWithError(res, http.StatusBadRequest, "Bad Request")
		return
	}

	// Create user
	userModel, err = app.repo.CreateUser(req.Context(), createUserRequest.Email, createUserRequest.Password, createUserRequest.FirstName, createUserRequest.LastName)
	if err != nil {
		app.respondWithError(res, http.StatusInternalServerError, "Critical server error happened")
		return
	}

	app.respondWithJwtToken(res, userModel.Email)

}

func (app *App) Login(res http.ResponseWriter, req *http.Request) {

	var credentials shared.Credentials

	// Decode Request
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		log.Error("Error while decoding request body: ", err.Error())
		app.respondWithError(res, http.StatusBadRequest, "Cannot decode request body")
		return
	}

	// Validate Request
	err := credentials.Validate()
	if err != nil {
		app.respondWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	// Get User
	userModel, err := app.repo.GetUserByEmail(req.Context(), credentials.Email)
	if err != nil {
		app.respondWithError(res, http.StatusBadRequest, "Bad Request")
		return
	}

	app.respondWithJwtToken(res, userModel.Email)

}

func (app *App) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	// Fetch users from DB
	userModels, err := app.repo.GetAllUsers(req.Context())
	if err != nil {
		log.Error("Error while fetching all users: ", err.Error())
		app.respondWithError(res, http.StatusInternalServerError, "Critical server error happened")
		return
	}

	// Convert DB model to API model
	var users []shared.User
	var resp shared.UserResponse

	for _, user := range *userModels {
		users = append(users, shared.User{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}
	resp = shared.UserResponse{Users: users}

	app.respondWithJSON(res, http.StatusOK, resp)
}

func (app *App) UpdateUser(res http.ResponseWriter, req *http.Request) {
	var updateUserRequest shared.UpdateUserRequest

	// Decode Request
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&updateUserRequest); err != nil {
		log.Error("Error while decoding request body: ", err.Error())
		app.respondWithError(res, http.StatusBadRequest, "Cannot decode request body")
		return
	}

	// Validate Request
	err := updateUserRequest.Validate()
	if err != nil {
		app.respondWithError(res, http.StatusBadRequest, err.Error())
		return
	}

	userEmail := req.Context().Value("email")
	if userEmail == nil {
		app.respondWithError(res, http.StatusInternalServerError, "Critical server error happened")
		return
	}

	// Update user info
	_, err = app.repo.UpdateUserInfo(req.Context(), userEmail.(string), updateUserRequest.FirstName, updateUserRequest.LastName)
	if err != nil {
		app.respondWithError(res, http.StatusInternalServerError, "Critical server error happened")
		return
	}

	app.respondWithJSON(res, http.StatusOK, nil)
}

func (app *App) respondWithJwtToken(res http.ResponseWriter, email string) {
	// Generate jwt token
	tokenString, err := jwt.Generate(email)
	if err != nil {
		app.respondWithError(res, http.StatusInternalServerError, "Critical server error happened")
		return
	}

	app.respondWithJSON(res, http.StatusOK, shared.TokenResponse{Token: tokenString})
}

func (app *App) respondWithError(res http.ResponseWriter, code int, message string) {
	app.respondWithJSON(res, code, map[string]string{"error": message})
}

func (app *App) respondWithJSON(res http.ResponseWriter, code int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)

	if payload != nil {
		response, _ := json.Marshal(payload)
		res.Write(response)
	}
}
