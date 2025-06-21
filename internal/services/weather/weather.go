package weather

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/errors"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container/services"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type weatherService struct {
	clients  *interfaces.ClientsContainer
	services *services.ServicesContainer
}

func NewWeatherService(clients *interfaces.ClientsContainer, services *services.ServicesContainer) interfaces.WeatherServiceInterface {
	return &weatherService{
		clients:  clients,
		services: services,
	}
}

func (s *weatherService) GetWeather(ctx context.Context, city string) (*structs.Weather, error) {
	res, err := s.clients.WeatherApi.GetWeather(ctx, city)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.ErrWeatherNotFound
	}

	return structs.NewWeatherFromCelsius(res.Current.TempC), nil
}
