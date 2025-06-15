package weather

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type weatherService struct {
	clients *interfaces.ClientsContainer
}

func NewWeatherService(clients *interfaces.ClientsContainer) interfaces.WeatherServiceInterface {
	return &weatherService{
		clients: clients,
	}
}

func (s *weatherService) GetWeather(ctx context.Context, city string) (*structs.WeatherResponse, error) {
	return s.clients.WeatherApi.GetWeather(ctx, city)
}
