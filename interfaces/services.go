package interfaces

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type CepServiceInterface interface {
	GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error)
	GetCepWeatherInfo(ctx context.Context, cep string) (*structs.Weather, error)
}

type WeatherServiceInterface interface {
	GetWeather(ctx context.Context, city string) (*structs.Weather, error)
}
