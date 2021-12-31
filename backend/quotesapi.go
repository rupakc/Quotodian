package main

import (
	"log"
	"net/http"
  quotesRouter "backend/routing"
)

func main() {
	router := quotesRouter.NewRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}
