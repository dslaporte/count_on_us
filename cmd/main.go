package main

import (
	"count_on_us/configs"
	factoryConnection "count_on_us/internal/infrastructure/db/sqlx/connection"
	"count_on_us/internal/infrastructure/webserver"
	ws "count_on_us/internal/infrastructure/webserver/go-chi"
	handlers "count_on_us/internal/infrastructure/webserver/go-chi/handlers"
	repositories "count_on_us/internal/repositories/account"
	validator "count_on_us/pkg/validators"
)

func main() {
	dbConnection, err := factoryConnection.NewDatabaseFactory(&configs.Get().DatabaseConfig)
	validator := validator.NewValidator()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()
	//create webserver
	goChiWebServer := ws.NewGoChiWebServer(configs.Get().WebServer.Port)

	//create repositories
	accountRepository := repositories.NewAccountRepository(dbConnection)

	//create handlers
	accountHandler := handlers.NewWebAccountHandler(*accountRepository, validator)

	//register handlers
	goChiWebServer.AddHandler(webserver.POST, "/accounts", accountHandler.Create)
	goChiWebServer.AddHandler(webserver.GET, "/accounts/{id}", accountHandler.Get)
	goChiWebServer.AddHandler(webserver.GET, "/accounts", accountHandler.List)
	goChiWebServer.AddHandler(webserver.PUT, "/accounts/{id}", accountHandler.Update)

	server := webserver.NewWebServerStarter(goChiWebServer)
	server.Start()
}
