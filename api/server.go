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
	a.Router.HandleFunc("/v1/api/users", listUsers).Methods("GET")
	a.Router.HandleFunc("/v1/api/users/{id}", retrieveUser).Methods("GET")
	a.Router.HandleFunc("/v1/api/users/register", createUser).Methods("POST")
	a.Router.HandleFunc("/v1/api/users/register/{id}", deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/v1/api/users/login", loginUser).Methods("POST")
	a.Router.HandleFunc("/v1/api/users/login/{id}", logoutUser).Methods("POST")

	return nil
}

// Initialize starts up the API connection
func (a *Server) Initialize(host string, port int32, user string, pw string, dbname string) {
	// Login to db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pw, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check connection to db
	err = db.Ping()
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
