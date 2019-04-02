package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/anthonydevelops/team-gen/api"
	"github.com/spf13/viper"
)

var a api.Server

func TestMain(m *testing.M) {
	a = api.Server{}
	a.Initialize(viper.GetString("DB.host"),
		viper.GetInt32("DB.port"),
		viper.GetString("DB.user"),
		viper.GetString("DB.password"),
		viper.GetString("DB.dbname"))

	timer := time.NewTimer(3 * time.Second)
	go func() {
		a.Run(viper.GetString("Server.port"))
		<-timer.C
		fmt.Println("Timer expired")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("Timer stopped")
	}
}
