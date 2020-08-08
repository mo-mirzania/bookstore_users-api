package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo-mirzania/bookstore_users-api/domain/users"
	"github.com/mo-mirzania/bookstore_users-api/services"
	"github.com/mo-mirzania/bookstore_users-api/utils/errors"
)

// GetUser func
func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "Implement me!")
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
	if err != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	fmt.Println(result)
	c.String(http.StatusOK, user.FirstName)
}

// SearchUser func
func SearchUser(c *gin.Context) {

}
