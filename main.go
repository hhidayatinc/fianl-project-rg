package main

import (
	"final-project/api"
	"final-project/repository"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main(){
	db, err := sql.Open("sqlite3", "./beasiswa.db")
	if err != nil {
		panic(err)
	}

	siswaRepo := repository.NewSiswaRepository(db)
	beasiswaRepo := repository.NewBeasiswaRepository(db)
	mainApi := api.NewApi(*siswaRepo, *beasiswaRepo)
	mainApi.Start() 
}