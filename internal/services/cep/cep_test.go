package cep

import (
	"context"
	"testing"

	"github.com/amandavmanduca/fullcycle-gcr/errors"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
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
		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi: viaCepApiMock,
		})
		address, err := service.GetAddress(ctx, cep)
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

		weatherApiMock := mocks.NewMockWeatherApiInterface(t)
		weatherApiMock.EXPECT().GetWeather(ctx, "Test City").Return(&structs.WeatherResponse{
			Current: structs.CurrentWeather{
				TempC: 25.0,
			},
		}, nil).Times(1)

		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi:  viaCepApiMock,
			WeatherApi: weatherApiMock,
		})

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

		weatherApiMock := mocks.NewMockWeatherApiInterface(t)

		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi:  viaCepApiMock,
			WeatherApi: weatherApiMock,
		})

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

		weatherApiMock := mocks.NewMockWeatherApiInterface(t)
		weatherApiMock.EXPECT().GetWeather(ctx, "Test City").Return(nil, nil).Times(1)

		service := NewCepService(&interfaces.ClientsContainer{
			ViaCepApi:  viaCepApiMock,
			WeatherApi: weatherApiMock,
		})

		weather, err := service.GetCepWeatherInfo(ctx, cep)
		assert.Nil(t, weather)
		assert.NotNil(t, err)
		assert.Equal(t, errors.ErrWeatherNotFound, err)
	})
}
