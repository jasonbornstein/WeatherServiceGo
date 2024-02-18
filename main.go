package main

import (
	"fmt"
	"log"

	"github.com/jasonbornstein/WeatherServiceGo/api"
	"github.com/jasonbornstein/WeatherServiceGo/config"
)

func main() {
	conf := config.LoadConfig()
	router := api.SetupRouter(conf)

	// Run the router
	err := router.Run(fmt.Sprintf("%s:%s", "localhost", "8080"))
	if err != nil {
		log.Fatalf("Error running server: %s\n", err)
		return
	}
}
