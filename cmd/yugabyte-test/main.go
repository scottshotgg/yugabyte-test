package main

import (
	"fmt"
	"time"

	"github.com/scottshotgg/yugabyte-test/db"
	"github.com/scottshotgg/yugabyte-test/pkg/http/rest"
)

const (
	host     = "yb-tserver-n1"
	port     = 5433
	user     = "yugabyte"
	password = "yugabyte"
	dbname   = "yugabyte"
)

func main() {
	fmt.Println("ya hai")

	// Just wait a little bit for Yugabyte to start in docker-compose
	time.Sleep(2 * time.Second)

	var dbb, err = db.NewYB(port, host, user, password, dbname)
	if err != nil {
		panic(err)
	}

	api := rest.New(dbb)
	api.Addr = "0.0.0.0:9090"
	api.ListenAndServe()

}
