package main

import (
	"net/http"
	"os"

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
	)

	var (
		docenteService = docente.NerService(docenteRepository)
	)

	r := chi.NewRouter()

	// r.Use(helper.GetCors().Handler)
	r.Mount("/docente", docente.MakeHTTPSHandler(docenteService))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	http.ListenAndServe(":"+PORT, r)
}
