package app

import (
	"github.com/mo-mirzania/bookstore_users-api/controllers/ping"
	"github.com/mo-mirzania/bookstore_users-api/controllers/users"
)

// Routes func
func Routes() {
	r.GET("/ping", ping.Ping) // We are not calling the Ping() function. We just specify which function is in charge of.
	//r.GET("/users/:user_id", users.GetUser)
	r.POST("/users", users.CreateUser) // We must not use /users/ !!!
	//r.GET("/users/search", controllers.SearchUser)
}
