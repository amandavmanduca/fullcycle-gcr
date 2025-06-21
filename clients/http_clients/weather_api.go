package http_clients

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type weatherApiClient struct {
	httpClient interfaces.HttpClientInterface
	apiKey     string
}

func NewWeatherApiClient(httpClient interfaces.HttpClientInterface, apiKey string) interfaces.WeatherApiInterface {
	return &weatherApiClient{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}

func (c *weatherApiClient) GetWeather(ctx context.Context, city string) (*structs.WeatherResponse, error) {
	params := map[string]string{
		"key": c.apiKey,
		"q":   city,
	}
	resp, err := c.httpClient.Get(ctx, "/current.json", params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("ERROR_GETTING_WEATHER")
	}

	var response structs.WeatherResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return &response, nil
	}
	return nil, errors.New("ERROR_GETTING_WEATHER")
}
