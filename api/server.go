// Package api provides endpoints & a connection to Postgres

//go:generate protoc --proto_path=model --proto_path=$GOPATH/src --gogo_out=model ./model/user.proto

package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server struct holds the backend API struct info
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

// InitializeRoutes sets & calls handlers
func (a *Server) InitializeRoutes() error {
	// Users route handles & endpoints
	a.Router.HandleFunc("/v1/api/users", a.ListUsers).Methods("GET")
	a.Router.HandleFunc("/v1/api/users/{id}", a.RetrieveUser).Methods("GET")
	a.Router.HandleFunc("/v1/api/users/register", a.CreateUser).Methods("POST")
	a.Router.HandleFunc("/v1/api/users/register/{id}", a.DeleteUser).Methods("DELETE")
	a.Router.HandleFunc("/v1/api/users/login", a.LoginUser).Methods("POST")
	a.Router.HandleFunc("/v1/api/users/login/{id}", a.LogoutUser).Methods("POST")

	return nil
}

// Initialize starts up the API connection
func (a *Server) Initialize(host string, port int32, user string, pw string, dbname string) {
	var err error

	// Login to db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pw, dbname)
	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer a.DB.Close()

	// Check connection to db
	err = a.DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("PostgreSQL successfully connected!")

	// Initialize routes
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

// Run exposes backend on port 8000
func (a *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
