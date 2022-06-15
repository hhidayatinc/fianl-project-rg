package main

import (
	"database/sql"

	"final-project/api"
	"final-project/db/migration"
	"final-project/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./beasiswa.db")
	if err != nil {
		panic(err)
	}
	migration.Generate(db)

	siswaRepo := repository.NewSiswaRepository(db)
	beasiswaRepo := repository.NewBeasiswaRepository(db)
	pendaftaranRepo := repository.NewPendaftaranRepository(db)
	mainApi := api.NewApi(*siswaRepo, *beasiswaRepo, *pendaftaranRepo)
	mainApi.Start()
}
