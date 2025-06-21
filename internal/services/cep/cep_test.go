package cep

import (
	"context"
	"testing"

	"github.com/amandavmanduca/fullcycle-gcr/errors"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container/services"
	"github.com/amandavmanduca/fullcycle-gcr/mocks"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {
	t.Run("should call cep client", func(t *testing.T) {
		ctx := context.Background()
		cep := "12345678"
		viaCepApiMock := mocks.NewMockViaCepApiInterface(t)
		viaCepApiMock.EXPECT().GetAddress(ctx, cep).Return(&structs.AddressResponse{
			Address: structs.Address{
				City:         "Test City",
				State:        "TS",
				Cep:          cep,
				Neighborhood: "Test Neighborhood",
				Street:       "Test Street",
			},
			Origin: structs.VIA_CEP,
		}, nil).Times(1)

		cepService := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi: viaCepApiMock,
		}, &services.ServicesContainer{})

		address, err := cepService.GetAddress(ctx, cep)
		assert.Nil(t, err)
		assert.Equal(t, "Test City", address.Address.City)
	})
}

func TestGetCepWeatherInfo(t *testing.T) {
	t.Run("should return weather info for a valid cep", func(t *testing.T) {
		ctx := context.Background()
		cep := "12345678"
		viaCepApiMock := mocks.NewMockViaCepApiInterface(t)
		viaCepApiMock.EXPECT().GetAddress(ctx, cep).Return(&structs.AddressResponse{
			Address: structs.Address{
				City:         "Test City",
				State:        "TS",
				Cep:          cep,
				Neighborhood: "Test Neighborhood",
				Street:       "Test Street",
			},
			Origin: structs.VIA_CEP,
		}, nil).Times(1)

		servicesContainer := &services.ServicesContainer{}
		mockWeatherService := mocks.NewMockWeatherServiceInterface(t)
		mockWeatherService.EXPECT().GetWeather(ctx, "Test City").Return(&structs.Weather{
			TempC: 25.0,
			TempF: 77.0,
			TempK: 298.0,
		}, nil).Times(1)

		servicesContainer.WeatherService = mockWeatherService
		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi: viaCepApiMock,
		}, servicesContainer)

		weather, err := service.GetCepWeatherInfo(ctx, cep)
		assert.Nil(t, err)
		assert.Equal(t, 25.0, weather.TempC)
		assert.Equal(t, 77.0, weather.TempF)
		assert.Equal(t, 298.0, weather.TempK)
	})
	t.Run("should return error for invalid city", func(t *testing.T) {
		ctx := context.Background()
		cep := "12345678"
		viaCepApiMock := mocks.NewMockViaCepApiInterface(t)
		viaCepApiMock.EXPECT().GetAddress(ctx, cep).Return(&structs.AddressResponse{
			Address: structs.Address{
				City:         "",
				State:        "TS",
				Cep:          cep,
				Neighborhood: "Test Neighborhood",
				Street:       "Test Street",
			},
			Origin: structs.VIA_CEP,
		}, nil).Times(1)

		servicesContainer := &services.ServicesContainer{}

		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi: viaCepApiMock,
		}, servicesContainer)

		weather, err := service.GetCepWeatherInfo(ctx, cep)
		assert.NotNil(t, err)
		assert.Nil(t, weather)
		assert.Equal(t, errors.ErrCannotFindZipcode, err)
	})
	t.Run("should return error if weather is not found", func(t *testing.T) {
		ctx := context.Background()
		cep := "12345678"
		viaCepApiMock := mocks.NewMockViaCepApiInterface(t)
		viaCepApiMock.EXPECT().GetAddress(ctx, cep).Return(&structs.AddressResponse{
			Address: structs.Address{
				City:         "Test City",
				State:        "TS",
				Cep:          cep,
				Neighborhood: "Test Neighborhood",
				Street:       "Test Street",
			},
			Origin: structs.VIA_CEP,
		}, nil).Times(1)

		servicesContainer := &services.ServicesContainer{}
		mockWeatherService := mocks.NewMockWeatherServiceInterface(t)
		mockWeatherService.EXPECT().GetWeather(ctx, "Test City").Return(nil, errors.ErrWeatherNotFound).Times(1)

		servicesContainer.WeatherService = mockWeatherService
		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi: viaCepApiMock,
		}, servicesContainer)

		weather, err := service.GetCepWeatherInfo(ctx, cep)
		assert.Nil(t, weather)
		assert.NotNil(t, err)
		assert.Equal(t, errors.ErrWeatherNotFound, err)
	})
}
