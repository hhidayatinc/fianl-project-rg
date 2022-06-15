package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Meta struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
type PendaftaranReponse struct {
	Meta    Meta   `json:"content"`
	Message string `json:"message"`
}

func (a *API) RegisterBeasiswa(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	ctx := r.Context()
	ctxVal := ctx.Value("email")
	emailPendaftar := fmt.Sprintf("%s", ctxVal)

	siswa, err := a.siswaRepo.GetByEmail(emailPendaftar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranReponse{
			Meta: Meta{
				Status: "error",
				Code:   http.StatusBadRequest,
			},
			Message: "siswa not found",
		})
		return
	}

	beasiswa, err := a.beasiswaRepo.GetById(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranReponse{
			Meta: Meta{
				Status: "error",
				Code:   http.StatusBadRequest,
			},
			Message: "beasiswa not found",
		})
		return
	}

	idBeasiswa := beasiswa.Id
	idSiswa := siswa.Id

	_, err = a.pendaftaranRepo.RegisterBeasiswa(int(idBeasiswa), int(idSiswa))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(PendaftaranReponse{
			Meta: Meta{
				Status: "error",
				Code:   http.StatusBadRequest,
			},
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(PendaftaranReponse{
		Meta: Meta{
			Status: "success",
			Code:   http.StatusOK,
		},
		Message: "Terdaftar",
	})
}
