package model 

import "golang.org/x/crypto/bcrypt"

// User to store user info
type User struct {
	Email     string `gorm:"primaryKey"`
	Password  string 
	FirstName string 
	LastName  string
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
