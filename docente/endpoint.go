package docente

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addDocenteRequest struct {
	Nombre          string
	ApellidoPat     string
	ApellidoMat     string
	Genero          string
	FechaNacimiento string
	Dni             string
}

func makeAddDocenteEndpoint(s Service) endpoint.Endpoint {
	addPersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addDocenteRequest)
		result, err := s.CrearDocente(&req)
		return result, err
	}
	return addPersonEndpoint
}
