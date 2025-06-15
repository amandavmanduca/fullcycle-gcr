package interfaces

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type CepServiceInterface interface {
	GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error)
}

type WeatherServiceInterface interface {
	GetWeather(ctx context.Context, city string) (*structs.WeatherResponse, error)
}
