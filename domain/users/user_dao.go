package users

import (
	"fmt"

	"github.com/moatazsalemVF/bookstore_user-api/datasources/mysqlds"
	"github.com/moatazsalemVF/bookstore_user-api/utils/datetime"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
	"github.com/moatazsalemVF/bookstore_user-api/utils/mysqlutils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?)"
	querySelectUserByID   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id = ?"
	queryDeleteUser       = "DELETE FROM users WHERE id = ?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?"
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
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {

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

	user.DateCreated = datetime.GetMysqlCurrentTimeUTC()
	user.Status = StatusActive

	stmt, err := mysqlds.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	stmtResult, mysqlErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

//FindUsersByStatus is a function to get users with specific status
func (user *User) FindUsersByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := mysqlds.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		u := User{}
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		users = append(users, u)
	}

	if len(users) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users found with the status '%s'", status))
	}

	return users, nil
}
