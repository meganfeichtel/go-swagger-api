// GoSwagger Example API Specification
//
// This is an example sepecification of how to use GoSwagger to docment an API written in GoLang
//
// There are currently no terms of service, so use at your owk risk.
//
// Terms of Service:
//
//     Schemes: https
//     Host: localhost
//     BasePath: /
//     Version: 1.1.0
//     Contact: Your Name Here <yname@redventures.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
