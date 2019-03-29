// Package routes provides endpoints & a connection to PostgreSQL

//go:generate protoc --proto_path=model --proto_path=$GOPATH/src --gogo_out=model ./model/user.proto

package routes

import (
	"github.com/gorilla/mux"
)

// NewRouter returns a new mux router w/ handlers
func NewRouter() *mux.Router {
	// Create new mux router
	r := mux.NewRouter()

	// List Users route
	r.HandleFunc("/v1/api/users", listUsers).Methods("GET")

	// Get User route
	r.HandleFunc("/v1/api/users/{id}", retrieveUser).Methods("GET")

	// Create & Delete User routes
	r.HandleFunc("/v1/api/users/register", createUser).Methods("POST")
	r.HandleFunc("/v1/api/users/register/{id}", deleteUser).Methods("DELETE")

	// Login & Logout User routes
	r.HandleFunc("/v1/api/users/login", loginUser).Methods("POST")
	r.HandleFunc("/v1/api/users/login/{id}", logoutUser).Methods("POST")

	return r
}
