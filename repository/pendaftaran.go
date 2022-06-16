package repository

import (
	"database/sql"
	"time"
)

type PendaftaranRepository struct {
	db *sql.DB
}

func NewPendaftaranRepository(db *sql.DB) *PendaftaranRepository {
	return &PendaftaranRepository{db}
}

func (r *PendaftaranRepository) RegisterBeasiswa(idBeasiswa, idSiswa int) (Pendaftaran, error) {
	t := time.Now()
	tanggal := t.Format("31-12-2006")
	status := "terdaftar"
	_, err := r.db.Exec("INSERT INTO pendaftaran (id_beasiswa, id_siswa, tanggal_daftar, status) VALUES (?, ?, ?, ?)", idBeasiswa, idSiswa, tanggal, status)
	if err != nil {
		return Pendaftaran{}, err
	}
	return Pendaftaran{}, nil
}

func (r *PendaftaranRepository) GetAllRegister() ([]Pendaftaran, error) {
	var p []Pendaftaran

	rows, err := r.db.Query("SELECT * FROM pendaftaran")
	if err != nil {
		return p, err
	}
	for rows.Next() {
		var pd Pendaftaran
		err := rows.Scan(&pd.Id, &pd.IdBeasiswa, &pd.IdSiswa, &pd.TanggalDaftar, &pd.Status)
		if err != nil {
			return p, err
		}
		p = append(p, pd)
	}
	return p, nil
}
