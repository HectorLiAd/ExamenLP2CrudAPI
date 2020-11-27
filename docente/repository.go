package docente

import "database/sql"

/*Repository permite interactuar con la BD*/
type Repository interface {
	CrearDocente(params *addDocenteRequest) (int, error)
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

func (r repository) CrearDocente(params *addDocenteRequest) (int, error) {
	const queryStr = `INSERT INTO PROFESOR (NOMBRE, APELLIDO_PATERNO, APELLIDO_MATERNO, GENERO, FECHA_NACIMIENTO, DNI)
	VALUES (?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(queryStr, params.Nombre, params.ApellidoPat,
		params.ApellidoMat, params.Genero, params.FechaNacimiento, params.Dni)
	if err != nil {
		return -1, err
	}
	id, er := result.LastInsertId()
	return int(id), er
}
