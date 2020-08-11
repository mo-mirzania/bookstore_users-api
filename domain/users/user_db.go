package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mo-mirzania/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	queryReadUser   = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id = ?;"
	queryDeleteUser = "DELETE FROM users WHERE id = ?;"
)

var (
	// UsersDB Connection
	UsersDB       *sql.DB
	mysqlUsername = os.Getenv("MYSQL_USERNAME")
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlHost     = os.Getenv("MYSQL_HOST")
	mysqlPort     = os.Getenv("MYSQL_PORT")
	mysqlSchema   = os.Getenv("MYSQL_SCHEMA")
)

// Get func
func (user *User) Get() *errors.RestErr {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema)
	UsersDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	err = UsersDB.Ping()
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Println("Connected to DB!")
	}
	defer UsersDB.Close()
	stmt, err := UsersDB.Prepare(queryReadUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreate)
	if err != nil {
		fmt.Println(err.Error())
		return errors.NewInternalServerError("Error when trying to get user")
	}
	return nil
}

// Save func
func (user *User) Save() *errors.RestErr {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema)
	UsersDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	err = UsersDB.Ping()
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Println("Connected to DB!")
	}
	defer UsersDB.Close() // It is very important actually!!!!!!
	stmt, err := UsersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	now := time.Now()
	user.DateCreate = now.Format("01-02-2006")

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreate)
	if err != nil {
		return errors.NewInternalServerError("Error when trying to save user!")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error when trying to get last insert ID")
	}
	user.ID = userID
	return nil
}

// Update func
func (user *User) Update() *errors.RestErr {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema)
	UsersDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	err = UsersDB.Ping()
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Println("Connected to DB!")
	}
	defer UsersDB.Close() // It is very important actually!!!!!!
	stmt, err := UsersDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError("Cannot update!")
	}
	return nil
}

// Delete func
func (user *User) Delete() *errors.RestErr {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlSchema)
	UsersDB, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	err = UsersDB.Ping()
	if err != nil {
		log.Panic(err.Error())
	} else {
		fmt.Println("Connected to DB!")
	}
	defer UsersDB.Close() // It is very important actually!!!!!!
	stmt, err := UsersDB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return errors.NewInternalServerError("Cannot update!")
	}

	return nil
}
