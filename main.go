package main

import (
	"gotest/src/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Serving at 8080")
	log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
