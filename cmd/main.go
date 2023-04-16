package main

import (
	"count_on_us/config"
	factoryConnection "count_on_us/internal/infrastructure/db/sqlx/connection"
	"count_on_us/internal/infrastructure/webserver"
	gochiws "count_on_us/internal/infrastructure/webserver/go-chi"
	repositories "count_on_us/internal/repositories/account"
	validator "count_on_us/pkg/validators"
)

func main() {
	dbConnection, err := factoryConnection.NewDatabaseFactory(&config.Get().DatabaseConfig)
	validator := validator.NewValidator()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()
	//create webserver
	goChiWebServer := gochiws.NewGoChiWebServer(config.Get().WebServer.Port)

	//create repositories
	accountRepository := repositories.NewAccountRepository(dbConnection)

	//create handlers
	accountHandler := webserver.NewWebAccountHandler(*accountRepository, validator)

	//register handlers
	goChiWebServer.AddHandler("/accounts", accountHandler.Create)

	server := webserver.NewWebServerStarter(goChiWebServer)
	server.Start()
}
