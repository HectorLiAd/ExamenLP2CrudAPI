package docente

/*ResultData json para mostrar el resultado de salida de una operacion a la BD*/
type ResultData struct {
	Mensaje     string `json:"mansaje"`
	CodEfectado string `json:"codAfectado"`
}
