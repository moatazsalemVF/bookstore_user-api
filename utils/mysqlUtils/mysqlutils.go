package mysqlutils

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
)

const (
	indexUniqueError = 1062
	noRowsError      = "no rows in result set"
)

//IsEmptyResult checks whether a result set is empty or not
func IsEmptyResult(err error) bool {
	if strings.Contains(err.Error(), noRowsError) {
		return true
	}
	return false
}

//IsZeroRowsAffected checks if the query didn't affect any rows
func IsZeroRowsAffected(result *sql.Result) bool {
	return true
}

//HandleMysqlError handles errors returned from mysql
func HandleMysqlError(mysqlErr error) *errors.RestError {
	sqlErr, ok := mysqlErr.(*mysql.MySQLError)
	if !ok {
		return errors.NewInternalServerError(mysqlErr.Error())
	}

	switch sqlErr.Number {
	case indexUniqueError:
		return errors.NewBadRequestError(fmt.Sprintf("Unique Key Vaiolated: %v", sqlErr.Message))
	}

	return errors.NewInternalServerError(mysqlErr.Error())
}
