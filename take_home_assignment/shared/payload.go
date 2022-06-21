package shared

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

type CreateUserRequest struct {
	Email     string `json:"email" valid:"email, required"`
	Password  string `json:"password" valid:"type(string), minstringlength(8), required"`
	FirstName string `json:"firstName" valid:"type(string), minstringlength(1), required"`
	LastName  string `json:"lastName" valid:"type(string), minstringlength(1), required"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName" valid:"type(string), minstringlength(1), required"`
	LastName  string `json:"lastName" valid:"type(string), minstringlength(1), required"`
}

type Credentials struct {
	Email    string `json:"email" valid:"email, required"`
	Password string `json:"password" valid:"type(string), minstringlength(8), required"`
}

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserResponse struct {
	Users []User `json:"users"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (request *CreateUserRequest) Validate() error {
	// Trim request
	request.Email = strings.ToLower(strings.TrimSpace(request.Email))
	request.Password = strings.TrimSpace(request.Password)
	request.FirstName = strings.TrimSpace(request.FirstName)
	request.LastName = strings.TrimSpace(request.LastName)

	// Validate
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return err
	}

	return nil
}

func (request *UpdateUserRequest) Validate() error {
	// Trim request
	request.FirstName = strings.TrimSpace(request.FirstName)
	request.LastName = strings.TrimSpace(request.LastName)

	// Validate
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return err
	}

	return nil
}

func (request *Credentials) Validate() error {
	// Trim request
	request.Email = strings.ToLower(strings.TrimSpace(request.Email))
	request.Password = strings.TrimSpace(request.Password)

	// Validate
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return err
	}

	return nil
}
