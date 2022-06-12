package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type BeasiswaListErrorResponse struct {
	Error string `json:"error"`
}

type ListBeasiswa struct {
	Id string `json:"id"`
	Nama string `json:"nama"`
	JenisBeasiswa string `json:"jenis_beasiswa"`
	JenjangPendidikan string `json:"jenjang_pendidikan"`
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalSelesai string `json:"tanggal_selesai"`
}

type BeasiswaListSuccessResponse struct {
	Beasiswa []ListBeasiswa `json:"beasiswa"`
}

func (a *API) getBeasiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := BeasiswaListSuccessResponse{}
	response.Beasiswa = make([]ListBeasiswa, 0)

	beasiswa, err := a.beasiswaRepo.GetAll()
	defer func(){
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(BeasiswaListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	for _, b := range beasiswa {
		response.Beasiswa = append(response.Beasiswa, ListBeasiswa{
			Id: strconv.Itoa(int(b.Id)),
			Nama: b.Nama, 
			JenisBeasiswa: b.JenisBeasiswa, 
			JenjangPendidikan: b.JenjangPendidikan, 
			TanggalMulai: b.TanggalMulai, 
			TanggalSelesai: b.TanggalSelesai,
		})
	}
	
	encoder.Encode(response)
}

func (a *API) getBeasiswaById(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	response := BeasiswaListSuccessResponse{}
	response.Beasiswa = make([]ListBeasiswa, 0)

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)

	beasiswa, err := a.beasiswaRepo.GetById(int64(idInt))
	defer func(){
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(BeasiswaListErrorResponse{Error: err.Error()})
		}
	}()
	if err != nil {
		return
	}
	response.Beasiswa = append(response.Beasiswa, ListBeasiswa{
		Id: strconv.Itoa(int(beasiswa.Id)),
		Nama: beasiswa.Nama,
		JenisBeasiswa: beasiswa.JenisBeasiswa,
		JenjangPendidikan: beasiswa.JenjangPendidikan,
		TanggalMulai: beasiswa.TanggalMulai,
		TanggalSelesai: beasiswa.TanggalSelesai,
	})
	encoder.Encode(response)
}