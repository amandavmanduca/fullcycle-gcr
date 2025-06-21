package services

import "github.com/amandavmanduca/fullcycle-gcr/interfaces"

type ServicesContainer struct {
	CepService     interfaces.CepServiceInterface
	WeatherService interfaces.WeatherServiceInterface
}
