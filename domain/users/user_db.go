package users

import (
	"github.com/mo-mirzania/bookstore_users-api/utils/errors"
)

var usersDB = make(map[int64]*User)

// Get func
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError("User not found!")
	}
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.ID = result.ID
	user.DateCreate = result.DateCreate
	return nil
}

// Save func
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {
		return errors.NewNotFoundError("User Already Exists!")
	}
	usersDB[user.ID] = user
	return nil
}
