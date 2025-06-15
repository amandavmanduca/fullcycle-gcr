package cep

import (
	"context"

	"github.com/amandavmanduca/fullcycle-gcr/interfaces"
	"github.com/amandavmanduca/fullcycle-gcr/structs"
)

type CepService struct {
	clients *interfaces.ClientsContainer
}

func NewCepService(clients *interfaces.ClientsContainer) interfaces.CepServiceInterface {
	return &CepService{
		clients: clients,
	}
}

func (s *CepService) GetAddress(ctx context.Context, cep string) (*structs.AddressResponse, error) {
	return s.clients.ViaCepApi.GetAddress(ctx, cep)
}
