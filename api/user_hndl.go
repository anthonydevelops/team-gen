package api

import (
	"fmt"
	"net/http"
)

// ListUsers retrieves all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("lists all users")
}

// RetrieveUser retrieves a specific user
func RetrieveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list a specified user")
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a new user")
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete a user")
}

// LoginUser logs in a user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login a user")
}

// LogoutUser logs out a user
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout a user")
}
