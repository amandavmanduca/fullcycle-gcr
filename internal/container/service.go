package container

import (
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/services/cep"
	"github.com/amandavmanduca/fullcycle-gcr/internal/services/weather"
)

type ServicesContainer struct {
	CepService     interfaces.CepServiceInterface
	WeatherService interfaces.WeatherServiceInterface
}

func NewServicesContainer(clients *interfaces.ClientsContainer) *ServicesContainer {
	return &ServicesContainer{
		CepService:     cep.NewCepService(clients),
		WeatherService: weather.NewWeatherService(clients),
	}
}
