// Backend API
// Main handles the execution of routing & db connection

package main

import (
	"fmt"

	"github.com/anthonydevelops/team-gen/api"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Init initializes parsing & reading of config file
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

// Main calls our server, connects to db, and initializes routes
func main() {
	a := api.Server{}
	a.Initialize(viper.GetString("DB.host"),
		viper.GetInt32("DB.port"),
		viper.GetString("DB.user"),
		viper.GetString("DB.password"),
		viper.GetString("DB.dbname"))

	a.Run(viper.GetString("Server.port"))
}
