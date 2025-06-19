package structs

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestNewWeatherFromCelsius(t *testing.T) {
	t.Run("should create a Weather struct with correct temperature conversions", func(t *testing.T) {
		tempC := 25.0
		expectedTempF := 77.0
		expectedTempK := 298.0

		weather := NewWeatherFromCelsius(tempC)

		assert.Equal(t, tempC, weather.TempC)
		assert.Equal(t, expectedTempF, weather.TempF)
		assert.Equal(t, expectedTempK, weather.TempK)
	})
	t.Run("should create a Weather struct with correct temperature conversions", func(t *testing.T) {
		tempC := 0.0
		expectedTempF := 32.0
		expectedTempK := 273.0

		weather := NewWeatherFromCelsius(tempC)

		assert.Equal(t, tempC, weather.TempC)
		assert.Equal(t, expectedTempF, weather.TempF)
		assert.Equal(t, expectedTempK, weather.TempK)
	})
	t.Run("should create a Weather struct with correct temperature conversions", func(t *testing.T) {
		tempC := -10.2
		expectedTempF := 13.64
		expectedTempK := 262.8

		weather := NewWeatherFromCelsius(tempC)

		assert.Equal(t, tempC, weather.TempC)
		assert.Equal(t, expectedTempF, weather.TempF)
		assert.Equal(t, expectedTempK, weather.TempK)
	})
}
