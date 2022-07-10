package main

import (
	"project3/eventapp/app/factory"
	"project3/eventapp/app/migration"
	"project3/eventapp/middlewares"
	"project3/eventapp/routes"

	"project3/eventapp/app/config"
)

func main() {
	// connection database
	dbConn := config.InitDB()
	// migration table
	migration.Migration(dbConn)
	// routes
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)

	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
