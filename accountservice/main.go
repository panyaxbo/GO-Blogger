package main

import (
	"fmt"

	"github.com/panyaxbo/goblogger/accountservice/dbclient"

	"github.com/panyaxbo/goblogger/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767") //NEW

}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
