package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		// ctx  context.Context = context.Background()
		httpport  string = DefaultGetenv("HTTP_PORT", "80")
		sqlhost   string = DefaultGetenv("SQL_HOST", "localhost")
		sqluser   string = DefaultGetenv("SQL_HOST", "root")
		sqlpass   string = DefaultGetenv("SQL_HOST", "test")
		sqldbname string = DefaultGetenv("SQL_HOST", "testtest")
	)

	// Registering HTTP handlers
	registerHandlers()

	// Configuring Database
	err := ConfigureSQL(DBInfo{
		Host: sqlhost,
		Port: 3306,
		User: sqluser,
		Pass: sqlpass,
		Name: sqldbname,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Run HTTP Server
	log.Printf("Start listening to HTTP requests at port [%s]", httpport)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpport), nil))
}
