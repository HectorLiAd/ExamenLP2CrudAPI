package curso

import (
	"database/sql"
)

/*Repository permite interactuar con la BD*/
type Repository interface {
	CrearCurso(params *addCursoRequest) (int, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository permite crear el repositoriio*/
func NewRepository(dataBaseConnection *sql.DB) Repository {
	return &repository{
		db: dataBaseConnection,
	}
}
func (repo repository) CrearCurso(params *addCursoRequest) (int, error) {
	const queryStr = `INSERT INTO CURSO ( PROFESOR_ID, NOMBRE, DESCRIPCION)
	VALUES (?, ?, ?)`
	result, err := repo.db.Exec(queryStr, params.DocenteID, params.Nombre, params.Descripcion)
	if err != nil {
		return -1, nil
	}

	idAfected, er := result.LastInsertId()

	return int(idAfected), er
}
