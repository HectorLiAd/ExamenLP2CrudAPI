package main

import (
	"net/http"
	"os"

	"github.com/HectorLiAd/ExamenLP2CrudAPI/curso"
	"github.com/HectorLiAd/ExamenLP2CrudAPI/database"
	"github.com/HectorLiAd/ExamenLP2CrudAPI/docente"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB()

	defer db.Close()

	var (
		docenteRepository = docente.NewRepository(db)
		cursoRepository   = curso.NewRepository(db)
	)

	var (
		docenteService = docente.NerService(docenteRepository)
		cursoService   = curso.NerService(cursoRepository)
	)

	r := chi.NewRouter()

	// r.Use(helper.GetCors().Handler)
	r.Mount("/docente", docente.MakeHTTPSHandler(docenteService))
	r.Mount("/curso", curso.MakeHTTPSHandler(cursoService))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	http.ListenAndServe(":"+PORT, r)

	// https://examenlp2.herokuapp.com/
}
