package services

import (
	"github.com/moatazsalemVF/bookstore_user-api/domain/users"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
)

//CreateUser is the service function to create a user
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser is the service function to get a user
func GetUser(id int64) (*users.User, *errors.RestError) {
	user := users.User{}
	user.ID = id
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
