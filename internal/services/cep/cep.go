package cep

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/errors"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type CepService struct {
	clients *interfaces.ClientsContainer
}

func NewCepService(clients *interfaces.ClientsContainer) interfaces.CepServiceInterface {
	return &CepService{
		clients: clients,
	}
}

func (s *CepService) GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error) {
	return s.clients.ViaCepApi.GetAddress(ctx, cep)
}

func (s *CepService) GetCepWeatherInfo(ctx context.Context, cep string) (*structs.Weather, error) {
	address, err := s.GetAddress(ctx, cep)
	if err != nil {
		return nil, err
	}
	if address.Address.City == "" {
		return nil, errors.ErrCannotFindZipcode
	}
	weatherResponse, err := s.clients.WeatherApi.GetWeather(ctx, address.Address.City)
	if err != nil {
		return nil, err
	}
	if weatherResponse == nil {
		return nil, errors.ErrWeatherNotFound
	}

	return structs.NewWeatherFromCelsius(weatherResponse.Current.TempC), nil
}
