package users

import (
	"github.com/moatazsalemVF/bookstore_user-api/datasources/mysql"
	"github.com/moatazsalemVF/bookstore_user-api/utils/datetime"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get is a function to retrieve user info from DB
func (user *User) Get() *errors.RestError {

	if err := mysql.Client.Ping(); err != nil {
		panic(err)
	}

	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError("User Not Found")
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

//Save is a function to save users in MySQL DB
func (user *User) Save() *errors.RestError {
	current := userDB[user.ID]
	if current != nil {
		return errors.NewBadRequestError("User Already Exists")
	}

	user.DateCreated = datetime.GetCurrentTimeUTC()
	userDB[user.ID] = user
	return nil
}
