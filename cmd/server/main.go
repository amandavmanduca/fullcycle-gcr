package main

import (
	"log"
	"net/http"
	"os"

	"github.com/amandavmanduca/fullcycle-gcr/clients"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container"
	"github.com/amandavmanduca/fullcycle-gcr/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientsConfig := clients.ClientsConfig{
		WeatherApiKey: os.Getenv("WEATHER_API_KEY"),
	}
	clients := clients.NewClientsContainer(clientsConfig)

	services := container.NewServicesContainer(&clients)
	handler := handlers.NewHandlerContainers(services)

	http.HandleFunc("/address-info", handler.CepHandler.GetAddressInfo)
	http.ListenAndServe(":8080", nil)
}
