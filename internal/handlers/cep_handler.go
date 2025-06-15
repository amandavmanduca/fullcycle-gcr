package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container"
)

type cepHandler struct {
	cepService     interfaces.CepServiceInterface
	weatherService interfaces.WeatherServiceInterface
}

func NewCepHandler(services container.ServicesContainer) cepHandler {
	return cepHandler{
		cepService:     services.CepService,
		weatherService: services.WeatherService,
	}
}

func (h *cepHandler) GetAddressInfo(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cep := r.URL.Query().Get("cep")
	if cep == "" || len(cep) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("invalid zipcode")
		return
	}
	address, err := h.cepService.GetAddress(ctx, cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if address.Address.City == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("can not find zipcode")
		return
	}
	weather, err := h.weatherService.GetWeather(ctx, address.Address.City)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	response := fmt.Sprintf("Temperatura em %s: %f°C / %f°F", address.Address.City, weather.Current.TempC, weather.Current.TempF)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
