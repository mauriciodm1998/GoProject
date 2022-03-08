package main

import (
	"API/src/config"
	"API/src/router"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	config.Load()

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
	fmt.Printf(config.ConnectionString)
}
