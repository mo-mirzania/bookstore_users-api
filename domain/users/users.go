package users

import (
	"strings"

	"github.com/mo-mirzania/bookstore_users-api/utils/errors"
)

// User struct
type User struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"fname"`
	LastName   string `json:"lname"`
	Email      string `json:"email"`
	DateCreate string `json:"date"`
}

// ValidateUser func
func (user *User) ValidateUser() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
