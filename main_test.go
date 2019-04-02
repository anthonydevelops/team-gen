package main_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/anthonydevelops/team-gen/api"
)

var a api.Server

func TestMain(m *testing.M) {
	a = api.Server{}
	a.Initialize("127.0.0.1",
		80,
		"postgres",
		"docker",
		"test")

	timer := time.NewTimer(3 * time.Second)
	go func() {
		a.Run("8000")
		<-timer.C
		fmt.Println("Timer expired")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("Timer stopped")
	}
}
