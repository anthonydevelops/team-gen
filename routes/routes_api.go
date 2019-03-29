// Package routes provides endpoints & a connection to Postgres

//go:generate protoc --proto_path=model --proto_path=$GOPATH/src --gogo_out=model ./model/user.proto

package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Login holds login info for psql
type Login struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

// NewRouter returns a new mux router w/ handlers
func NewRouter(conf string) (router *mux.Router, err error) {
	// Open keys file
	file, err := os.Open(conf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	// Map keys json to Login struct
	login := []Login{}
	keys, _ := ioutil.ReadAll(file)
	json.Unmarshal(keys, &login)

	// Create new mux router
	r := mux.NewRouter()

	return r, nil
}

// HandleEndpoints sets & calls handlers
func HandleEndpoints(r *mux.Router) error {
	// Users route handles & endpoints
	r.HandleFunc("/v1/api/users", listUsers).Methods("GET")
	r.HandleFunc("/v1/api/users/{id}", retrieveUser).Methods("GET")
	r.HandleFunc("/v1/api/users/register", createUser).Methods("POST")
	r.HandleFunc("/v1/api/users/register/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/v1/api/users/login", loginUser).Methods("POST")
	r.HandleFunc("/v1/api/users/login/{id}", logoutUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

	return nil
}
