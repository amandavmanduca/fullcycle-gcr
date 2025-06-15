package clients

import (
	"github.com/amandavmanduca/fullcycle-gcr/clients/http_clients"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
)

type ClientsConfig struct {
	WeatherApiKey string
}

func NewClientsContainer(config ClientsConfig) interfaces.ClientsContainer {
	viaCepApi := NewHttpClient("http://viacep.com.br/ws", nil)
	weatherApi := NewHttpClient("http://api.weatherapi.com/v1", nil)

	return interfaces.ClientsContainer{
		ViaCepApi:  http_clients.NewViaCepApiClient(viaCepApi),
		WeatherApi: http_clients.NewWeatherApiClient(weatherApi, config.WeatherApiKey),
	}

}
