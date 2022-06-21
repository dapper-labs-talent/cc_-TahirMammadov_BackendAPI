package tests

import (
	"context"

	"github.com/tmammado/take-home-assignment/model"
)

type MockUserRepo struct {
	user model.User
}

type MockUserRepoDefault struct {
	user model.User
}

// MockUserRepo
func (repo *MockUserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (repo *MockUserRepo) CreateUser(ctx context.Context, email, password, firstName, lastName string) (*model.User, error) {
	return &repo.user, nil
}

func (repo *MockUserRepo) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	return &[]model.User{repo.user}, nil
}

func (repo *MockUserRepo) UpdateUserInfo(ctx context.Context, email, firstName, lastName string) (*model.User, error) {
	return &repo.user, nil
}

// MockUserRepoDefault methods

func (repo *MockUserRepoDefault) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return &repo.user, nil
}

func (repo *MockUserRepoDefault) CreateUser(ctx context.Context, email, password, firstName, lastName string) (*model.User, error) {
	return &repo.user, nil
}

func (repo *MockUserRepoDefault) GetAllUsers(ctx context.Context) (*[]model.User, error) {
	return &[]model.User{repo.user}, nil
}

func (repo *MockUserRepoDefault) UpdateUserInfo(ctx context.Context, email, firstName, lastName string) (*model.User, error) {
	return &repo.user, nil
}
