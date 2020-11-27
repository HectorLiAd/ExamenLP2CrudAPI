package curso

import (
	"errors"
	"strconv"

	"github.com/HectorLiAd/ExamenLP2CrudAPI/models"
)

/*Service interface para los sercicios*/
type Service interface {
	CrearCurso(params *addCursoRequest) (*models.ResultData, error)
}

type service struct {
	repo Repository
}

/*NerService permite crear el sercicio*/
func NerService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s service) CrearCurso(params *addCursoRequest) (*models.ResultData, error) {
	result, err := s.repo.CrearCurso(params)
	if err != nil {
		return nil, err
	}
	if result == -1 {
		return nil, errors.New("Verifique de que el usuario exista y valida los datos")
	}
	resultado := &models.ResultData{
		Mensaje:     "Se guardo correctamente " + params.Nombre + " con el id " + strconv.Itoa(result),
		CodEfectado: strconv.Itoa(result),
	}
	return resultado, nil
}
