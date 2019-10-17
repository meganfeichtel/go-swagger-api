// GoSwagger Example API Specification for ToDo List
//
// This is an example sepecification of how to use GoSwagger to docment an API written in GoLang
// implemnting the API example setup here: https://thenewstack.io/make-a-restful-json-api-go/
//
// There are currently no terms of service, so use at your own risk.
//
// Terms of Service:
//
//     Schemes: https
//     Host: localhost
//     BasePath: /
//     Version: 1.1.0
//     Contact: meganfeichtel in an issue
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
