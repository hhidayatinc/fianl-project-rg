package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type SiswaListErrorResponse struct {
	Error string `json:"error"`
}

type ListSiswa struct {
	Id string `json:"id"`
	Nama string `json:"nama"`
	TempatLahir string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenjangPendidikan string `json:"jenjang_pendidikan"`
	Nik string `json:"nik"`
	Email string `json:"email"`
}

type SiswaListSuccessResponse struct {
	Siswa []ListSiswa `json:"siswa"`
}

func (a *API) getSiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := SiswaListSuccessResponse{}
	response.Siswa = make([]ListSiswa, 0)

	siswa, err := a.siswaRepo.GetAll()
	defer func(){
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(SiswaListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, s := range siswa {
		response.Siswa = append(response.Siswa, ListSiswa{
			Id :  strconv.Itoa(int(s.Id)),
			Nama: s.Nama,
			Nik: s.Nik,
			TempatLahir: s.TempatLahir,
			TanggalLahir: s.TanggalLahir,
			JenjangPendidikan: s.JenjangPendidikan,
			Email: s.Email,
		})
	}
	encoder.Encode(response)
}

func (a *API) getSiswaById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := SiswaListSuccessResponse{}
	response.Siswa = make([]ListSiswa, 0)

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	siswa, err := a.siswaRepo.GetById(int64(id))
	defer func(){
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(SiswaListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Siswa = append(response.Siswa, ListSiswa{
		Id :  strconv.Itoa(int(siswa.Id)),
		Nama: siswa.Nama,
		Nik: siswa.Nik,
		TempatLahir: siswa.TempatLahir,
		TanggalLahir: siswa.TanggalLahir,
		JenjangPendidikan: siswa.JenjangPendidikan,
		Email: siswa.Email,
	})
	encoder.Encode(response)
}
