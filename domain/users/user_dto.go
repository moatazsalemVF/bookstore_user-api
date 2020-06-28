package users

import (
	"strings"

	"github.com/moatazsalemVF/bookstore_user-api/utils/errors"
)

//User is the main domain
type User struct {
	ID          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	DateCreated string `json:"date_created,omitempty"`
}

//Validate is used to validate user struct
func (u *User) Validate() *errors.RestError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("Invalid User Email Address")
	}
	return nil
}
