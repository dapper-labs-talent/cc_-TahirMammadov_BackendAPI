package repository

import (
	"context"
	"time"

	"github.com/go-playground/log/v8"
	"gorm.io/gorm"

	"github.com/tmammado/take-home-assignment/model"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	userModel := model.User{Email: email}

	// Check if user exist in db
	result := repo.db.First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userModel, nil
}

func (repo *UserRepo) CreateUser(ctx context.Context, email, password, firstName, lastName string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	userModel := model.User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	// Hash User Password
	err := userModel.HashPassword()
	if err != nil {
		log.Error("Error while hashing user password: %s", err.Error())
		return nil, err
	}

	// Store user in db
	result := repo.db.Create(&userModel)
	if result.Error != nil {
		log.Error("Error while storing user in DB: %s", result.Error)
		return nil, result.Error
	}

	return &userModel, nil
}

func (repo *UserRepo) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var userModels []model.User

	// Query db to get all users
	result := repo.db.Find(&userModels)
	if result.Error != nil {
		log.Error("Error while fetching all users: %s", result.Error)
		return nil, result.Error
	}

	return &userModels, nil
}

func (repo *UserRepo) UpdateUserInfo(ctx context.Context, email, firstName, lastName string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	userModel := model.User{
		Email: email,
	}

	result := repo.db.First(&userModel)
	if result.Error != nil {
		log.Error("Error while fetching user: %s", result.Error)
		return nil, result.Error
	}

	// Update user
	userModel.FirstName = firstName
	userModel.LastName = lastName
	result = repo.db.Save(&userModel)
	if result.Error != nil {
		log.Debug("Error while saving user to DB: %s", result.Error)
		return nil, result.Error
	}

	return &userModel, nil
}
