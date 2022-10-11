package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/db/postgres"
	"github.com/MydroX/technical-test-fleurs-d-ici/pkg/env"
	"github.com/MydroX/technical-test-fleurs-d-ici/src/router"
)

func main() {
	env.Load()

	postgres.InitDB()

	r := router.New()

	log.Println(fmt.Sprintf("Server is starting on port %v", os.Getenv("PORT")))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), r))
}
