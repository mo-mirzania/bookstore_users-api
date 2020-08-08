package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mo-mirzania/bookstore_users-api/domain/users"
	"github.com/mo-mirzania/bookstore_users-api/services"
	"github.com/mo-mirzania/bookstore_users-api/utils/errors"
)

// GetUser func
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("Invalid User ID")
		c.JSON(err.Status, err)
		return
	}

	result, restErr := services.GetUser(userID)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// CreateUser func
func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.CreateUser(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.String(http.StatusOK, result.FirstName)
}

// SearchUser func
func SearchUser(c *gin.Context) {

}
