package routes

import (
	"fmt"
	"net/http"
)

// Retrieve all users
func listUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list all users")
}

// Retrieve a specific user
func retrieveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list a specified user")
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create a new user")
}

// Delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete a user")
}

// Login a user
func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login a user")
}

// Logout a user
func logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout a user")
}
