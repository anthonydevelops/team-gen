// Backend API
// Main handles the execution of handlers, generation, and connection to db

//go:generate protoc --proto_path=model --proto_path=$GOPATH/src --gogo_out=model ./model/user.proto

package main

import (
	"github.com/gorilla/mux"
)

func main() {
	//Initialize router
	r := mux.NewRouter()

	//Route handlers
	r.HandleFunc("/v1/api/user", getUser).Methods("GET")
}
