package services

import (
	"github.com/moatazsalemVF/bookstore_user-api/domain/users"
	"github.com/moatazsalemVF/bookstore_user-api/utils/crypto"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
)

//CreateUser is the service function to create a user
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Password = crypto.GetMD5(user.Password)

	if err := user.SaveOrUpdate(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser is the service function to update a user
func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestError) {

	saved := users.User{}
	saved.ID = user.ID
	if err := saved.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName == "" {
			user.FirstName = saved.FirstName
		}
		if user.LastName == "" {
			user.LastName = saved.LastName
		}
		if user.Email == "" {
			user.Email = saved.Email
		}
	}
	if err := user.SaveOrUpdate(); err != nil {
		return nil, err
	}
	return &user, nil
}

//DeleteUser is the service function to delete a user
func DeleteUser(id int64) *errors.RestError {
	user := users.User{}
	user.ID = id

	err := user.Get()
	if err != nil {
		return err
	}

	if err := user.Remove(); err != nil {
		return err
	}
	return nil
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

//FindUsersByStatus is the service function to get users with specific status
func FindUsersByStatus(status string) ([]users.User, *errors.RestError) {
	user := users.User{}
	return user.FindUsersByStatus(status)
}
