// Backend API
// Main handles the execution of routing & db connection

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/anthonydevelops/team-gen/routes"
)

var (
	conf string
)

func main() {
	RegisterFlags()

	// Initialize router
	r, err := routes.NewRouter(conf)
	if err != nil {
		return log.Fatal(err)
	}

	// Call endpoints
	err = routes.HandleEndpoints(r)
	if err != nil {
		return log.Fatal(err)
	}
}

// RegisterFlags registers startup flags
func RegisterFlags() {
	fmt.Println("Registering flags...")
	flag.StringVar(&conf, "keys-file", "keys.json", "A keys-file holds the db login information")

	flag.Parse()
}
