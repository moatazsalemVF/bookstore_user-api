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

	return &user, nil
}
