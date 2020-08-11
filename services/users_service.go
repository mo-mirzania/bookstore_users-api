package services

import (
	"github.com/mo-mirzania/bookstore_users-api/domain/users"
	"github.com/mo-mirzania/bookstore_users-api/utils/errors"
)

// GetUser func
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	err := result.Get()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUser func
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.ValidateUser()
	if err != nil {
		return nil, err
	}

	err = user.Save()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser func
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}
