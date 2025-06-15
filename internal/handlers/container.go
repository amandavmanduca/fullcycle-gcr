package handlers

import "github.com/amandavmanduca/fullcycle-gcr/internal/container"

type HandlerContainer struct {
	CepHandler cepHandler
}

func NewHandlerContainers(services *container.ServicesContainer) *HandlerContainer {
	return &HandlerContainer{
		CepHandler: NewCepHandler(*services),
	}
}
