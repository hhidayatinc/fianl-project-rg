package main

import (
	"final-project/api"
	"final-project/repository"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"final-project/db/migration"
)

func main(){
	db, err := sql.Open("sqlite3", "./beasiswa.db")
	if err != nil {
		panic(err)
	}
	migration.Generate(db)

	siswaRepo := repository.NewSiswaRepository(db)
	beasiswaRepo := repository.NewBeasiswaRepository(db)
	mainApi := api.NewApi(*siswaRepo, *beasiswaRepo)
	mainApi.Start() 
}