package main

import (
	"github.com/GermainSIGETY/golang-ddd-kata/internal/bootstrap"
	"github.com/rs/zerolog/log"
)

// @title Get Things Done Todos Rest API
// @version 1.0
// @description This is a golang project that serves as an example of hexagonal architecture

// @contact.email gsigety@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host staging.todo-company.co.uk
func main() {
	router := bootstrap.InitApp()
	if errRun := router.Run(); errRun != nil {
		log.Fatal().Err(errRun).Msg("Error during service instantiation")
	}
}
