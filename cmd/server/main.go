package main

import (
	"log"
	"net/http"
	"os"

	"github.com/amandavmanduca/fullcycle-gcr/clients"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container"
	"github.com/amandavmanduca/fullcycle-gcr/internal/handlers"
)

func main() {
	clientsConfig := clients.ClientsConfig{
		WeatherApiKey: os.Getenv("WEATHER_API_KEY"),
	}
	clients := clients.NewClientsContainer(clientsConfig)

	services := container.NewServicesContainer(&clients)
	handler := handlers.NewHandlerContainers(services)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Starting server on port", port)

	http.HandleFunc("/address-info", handler.CepHandler.GetAddressInfo)
	http.ListenAndServe(":"+port, nil)
}
