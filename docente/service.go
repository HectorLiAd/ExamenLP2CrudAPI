package docente

import (
	"strconv"

	"github.com/HectorLiAd/ExamenLP2CrudAPI/models"
)

/*Service interface para los sercicios*/
type Service interface {
	CrearDocente(params *addDocenteRequest) (*models.ResultData, error)
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

func (s service) CrearDocente(params *addDocenteRequest) (*models.ResultData, error) {
	result, err := s.repo.CrearDocente(params)
	if err != nil {
		return nil, err
	}
	resultado := &models.ResultData{
		Mensaje:     "Se guardo correctamente a " + params.Nombre + " con el id " + strconv.Itoa(result),
		CodEfectado: strconv.Itoa(result),
	}
	return resultado, nil
}
