package interfaces

import (
	"context"
	"net/http"

	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type HttpClientInterface interface {
	Get(ctx context.Context, path string) (*http.Response, error)
}

type ClientsContainer struct {
	ViaCepApi  ViaCepApiInterface
	WeatherApi WeatherApiInterface
}

type ViaCepApiInterface interface {
	GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error)
}

type WeatherApiInterface interface {
	GetWeather(ctx context.Context, city string) (*structs.WeatherResponse, error)
}
