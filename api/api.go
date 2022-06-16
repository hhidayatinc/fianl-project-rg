package api

import (
	"fmt"
	"net/http"

	"final-project/repository"
)

type API struct {
	mux             *http.ServeMux
	siswaRepo       repository.SiswaRepository
	beasiswaRepo    repository.BeasiswaRepository
	pendaftaranRepo repository.PendaftaranRepository
}

func NewApi(siswaRepo repository.SiswaRepository, beasiswaRepo repository.BeasiswaRepository, pendaftaranRepo repository.PendaftaranRepository) *API {
	mux := http.NewServeMux()

	api := &API{
		mux:             mux,
		siswaRepo:       siswaRepo,
		beasiswaRepo:    beasiswaRepo,
		pendaftaranRepo: pendaftaranRepo,
	}

	mux.Handle("/api/login", api.POST(http.HandlerFunc(api.login)))
	mux.Handle("/api/register", api.POST(http.HandlerFunc(api.register)))

	mux.Handle("/api/siswa/all", api.GET(http.HandlerFunc(api.getSiswa)))
	mux.Handle("/api/siswa", api.GET(http.HandlerFunc(api.getSiswaById)))

	mux.Handle("/api/beasiswa", api.GET(http.HandlerFunc(api.getBeasiswa)))
	mux.Handle("/api/beasiswa/", api.GET(http.HandlerFunc(api.getBeasiswaById)))
	mux.Handle("/api/beasiswa/register", api.AuthMiddleware(http.HandlerFunc(api.RegisterBeasiswa)))

	// mux.HandleFunc("/api/siswa", api.handleSiswa)
	// mux.HandleFunc("/api/siswa/{id}", api.getSiswaById)
	// mux.HandleFunc("/api/beasiswa", api.getBeasiswa)
	// mux.HandleFunc("/api/beasiswa/{id}", api.getBeasiswaById)
	// mux.HandleFunc("/api/login", api.login)
	// mux.HandleFunc("/api/register", api.register)
	// mux.HandleFunc("/api/logout", api.logout)
	// api.mux.HandleFunc("/beasiswa", api.getBeasiswa).Methods("GET")
	// api.mux.HandleFunc("/beasiswa/{id}", api.getBeasiswaById).Methods("GET")
	// api.mux.HandleFunc("/siswa", api.getSiswa).Methods("GET")
	// api.mux.HandleFunc("/siswa/{id}", api.getSiswaById).Methods("GET")
	// api.mux.HandleFunc("/login", api.login).Methods("POST")
	// api.mux.HandleFunc("/register", api.register).Methods("POST")
	// api.mux.HandleFunc("/logout", api.logout).Methods("POST")
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", api.Handler())
}
