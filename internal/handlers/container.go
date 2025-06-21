package handlers

import (
	"github.com/amandavmanduca/fullcycle-gcr/internal/container/services"
)

type HandlerContainer struct {
	CepHandler cepHandler
}

func NewHandlerContainers(services *services.ServicesContainer) *HandlerContainer {
	return &HandlerContainer{
		CepHandler: NewCepHandler(*services),
	}
}
