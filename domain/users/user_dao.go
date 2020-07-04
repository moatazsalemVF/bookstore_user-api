package users

import (
	"github.com/moatazsalemVF/bookstore_user-api/datasources/mysqlds"
	"github.com/moatazsalemVF/bookstore_user-api/utils/datetime"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
	"github.com/moatazsalemVF/bookstore_user-api/utils/mysqlutils"
)

const (
	queryInsertUser     = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?)"
	querySelectUserByID = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
	queryUpdateUser     = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id = ?"
	queryDeleteUser     = "DELETE FROM users WHERE id = ?"
)

//Remove is a function to save users in MySQL DB
func (user *User) Remove() *errors.RestError {
	stmt, err := mysqlds.Client.Prepare(queryDeleteUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, mysqlErr := stmt.Exec(user.ID)
	if mysqlErr != nil {
		return mysqlutils.HandleMysqlError(mysqlErr)
	}

	return nil
}

//Get is a function to retrieve user info from DB
func (user *User) Get() *errors.RestError {

	stmt, err := mysqlds.Client.Prepare(querySelectUserByID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {

		if mysqlutils.IsEmptyResult(err) {
			return errors.NewNotFoundError("User was not found")
		}

		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

//SaveOrUpdate is a function to save users in MySQL DB
func (user *User) SaveOrUpdate() *errors.RestError {
	if user.ID == 0 {
		return saveNewUser(user)
	}
	return updateUser(user)
}

func updateUser(user *User) *errors.RestError {

	stmt, err := mysqlds.Client.Prepare(queryUpdateUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, mysqlErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if mysqlErr != nil {
		return mysqlutils.HandleMysqlError(mysqlErr)
	}

	//get Updated record
	user.Get()

	return nil
}

func saveNewUser(user *User) *errors.RestError {
	user.DateCreated = datetime.GetCurrentTimeUTC()

	stmt, err := mysqlds.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	stmtResult, mysqlErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if mysqlErr != nil {
		return mysqlutils.HandleMysqlError(mysqlErr)
	}

	lastID, err := stmtResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	user.ID = lastID
	return nil
}
