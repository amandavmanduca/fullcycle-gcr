package weather

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

func TestGetWeather(t *testing.T) {
	t.Run("should return weather info for a valid cep", func(t *testing.T) {
		ctx := context.Background()
		city := "Test City"

		weatherApiMock := mocks.NewMockWeatherApiInterface(t)
		weatherApiMock.EXPECT().GetWeather(ctx, "Test City").Return(&structs.WeatherResponse{
			Current: structs.CurrentWeather{
				TempC: 25.0,
			},
		}, nil).Times(1)

		service := NewWeatherService(&interfaces.ClientsContainer{
			WeatherApi: weatherApiMock,
		}, &services.ServicesContainer{})

		weather, err := service.GetWeather(ctx, city)
		assert.Nil(t, err)
		assert.Equal(t, 25.0, weather.TempC)
		assert.Equal(t, 77.0, weather.TempF)
		assert.Equal(t, 298.0, weather.TempK)
	})
	t.Run("should return error for invalid city", func(t *testing.T) {
		ctx := context.Background()
		city := "Test City"

		weatherApiMock := mocks.NewMockWeatherApiInterface(t)
		weatherApiMock.EXPECT().GetWeather(ctx, "Test City").Return(nil, errors.ErrCannotFindZipcode).Times(1)

		service := NewWeatherService(&interfaces.ClientsContainer{
			WeatherApi: weatherApiMock,
		}, &services.ServicesContainer{})

		weather, err := service.GetWeather(ctx, city)
		assert.NotNil(t, err)
		assert.Nil(t, weather)
		assert.Equal(t, errors.ErrCannotFindZipcode, err)
	})
	t.Run("should return error if weather is not found", func(t *testing.T) {
		ctx := context.Background()
		city := "Test City"

		weatherApiMock := mocks.NewMockWeatherApiInterface(t)
		weatherApiMock.EXPECT().GetWeather(ctx, "Test City").Return(nil, nil).Times(1)

		service := NewWeatherService(&interfaces.ClientsContainer{
			WeatherApi: weatherApiMock,
		}, &services.ServicesContainer{})

		weather, err := service.GetWeather(ctx, city)
		assert.Nil(t, weather)
		assert.NotNil(t, err)
		assert.Equal(t, errors.ErrWeatherNotFound, err)
	})
}
