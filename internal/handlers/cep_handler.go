package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	localErrs "github.com/amandavmanduca/fullcycle-gcr/errors"
	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/internal/container"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type cepHandler struct {
	cepService interfaces.CepServiceInterface
}

func NewCepHandler(services container.ServicesContainer) cepHandler {
	return cepHandler{
		cepService: services.CepService,
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
	weather, err := h.cepService.GetCepWeatherInfo(ctx, cep)
	if err != nil {
		if errors.Is(err, localErrs.ErrCannotFindZipcode) || errors.Is(err, localErrs.ErrWeatherNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(
				structs.HttpResponse{
					Data:  nil,
					Error: err.Error(),
				})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			structs.HttpResponse{
				Data:  nil,
				Error: err.Error(),
			})
		return
	}

	mapResponse := map[string]float64{
		"temp_c": weather.TempC,
		"temp_f": weather.TempF,
		"temp_k": weather.TempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(structs.HttpResponse{
		Data:  mapResponse,
		Error: nil,
	})
}
