package curso

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addCursoRequest struct {
	DocenteID   int
	Nombre      string
	Descripcion string
}

func makeAddCursoEndpoint(s Service) endpoint.Endpoint {
	addCursoEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addCursoRequest)
		result, err := s.CrearCurso(&req)
		return result, err
	}
	return addCursoEndpoint
}
