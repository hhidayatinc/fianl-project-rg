package repository

import "database/sql"

type SiswaRepository struct {
	db *sql.DB
}

func NewSiswaRepository(db *sql.DB) *SiswaRepository {
	return &SiswaRepository{db: db}
}

func (r *SiswaRepository) GetAll() ([]Siswa, error) {
	var siswa []Siswa
	rows, err := r.db.Query("SELECT * FROM siswa")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var s Siswa
		err := rows.Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
		if err != nil {
			return nil, err
		}
		siswa = append(siswa, s)
	}
	return siswa, nil
}

func (r *SiswaRepository) GetById(id int64) (Siswa, error) {
	var s Siswa
	err := r.db.QueryRow("SELECT * FROM siswa WHERE id = ?", id).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *SiswaRepository) Register(nama string, password string, email string, jenjangPendidikan string, nik string, tanggalLahir string, tempatLahir string) (Siswa, error) {
	var s Siswa
	err := r.db.QueryRow("INSERT INTO siswa (nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING id, nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir", nama, password, email, jenjangPendidikan, nik, tanggalLahir, tempatLahir).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *SiswaRepository) Login(email string, password string) (Siswa, error) {
	var s Siswa
	err := r.db.QueryRow("SELECT * FROM siswa WHERE email = ? AND password = ?", email, password).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *SiswaRepository) Update(id int64, nama string, password string, email string, jenjangPendidikan string, nik string, tanggalLahir string, tempatLahir string) (Siswa, error) {
	var s Siswa
	err := r.db.QueryRow("UPDATE siswa SET nama = ?, password = ?, email = ?, jenjang_pendidikan = ?, nik = ?, tanggal_lahir = ?, tempat_lahir = ? WHERE id = ? RETURNING id, nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir", nama, password, email, jenjangPendidikan, nik, tanggalLahir, tempatLahir, id).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}

func (r *SiswaRepository) Logout(id int64) (Siswa, error) {
	var s Siswa
	err := r.db.QueryRow("UPDATE siswa SET password = ? WHERE id = ? RETURNING id, nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir", "", id).Scan(&s.Id, &s.Nama, &s.Password, &s.Email, &s.JenjangPendidikan, &s.Nik, &s.TanggalLahir, &s.TempatLahir)
	if err != nil {
		return s, err
	}
	return s, nil
}
