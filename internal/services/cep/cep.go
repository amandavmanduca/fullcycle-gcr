package cep

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/errors"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container/services"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type CepService struct {
	clients  *interfaces.ClientsContainer
	services *services.ServicesContainer
}

func NewCepService(clients *interfaces.ClientsContainer, services *services.ServicesContainer) interfaces.CepServiceInterface {
	return &CepService{
		clients:  clients,
		services: services,
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
	return s.services.WeatherService.GetWeather(ctx, address.Address.City)
}
