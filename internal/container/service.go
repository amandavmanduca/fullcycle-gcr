package container

import (
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container/services"
	"github.com/amandavmanduca/fullcycle-gcr/internal/services/cep"
	"github.com/amandavmanduca/fullcycle-gcr/internal/services/weather"
)

func NewServicesContainer(clients *interfaces.ClientsContainer) *services.ServicesContainer {
	container := &services.ServicesContainer{}
	container.WeatherService = weather.NewWeatherService(clients, container)
	container.CepService = cep.NewCepService(clients, container)
	return container
}
