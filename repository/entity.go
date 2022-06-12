package repository

type Beasiswa struct {
	Id          int64     `db:"id" json:"id"`
	Nama        string    `db:"nama" json:"nama"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	JenisBeasiswa string    `db:"jenis_beasiswa" json:"jenis_beasiswa"`
	JenjangPendidikan string `db:"jenjang_pendidikan" json:"jenjang_pendidikan"`
	TanggalMulai string `db:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai string `db:"tanggal_selesai" json:"tanggal_selesai"`
}

type Siswa struct {
	Id          int64     `db:"id" json:"id"`
	Nama        string    `db:"nama" json:"nama"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	JenjangPendidikan string `db:"jenjang_pendidikan" json:"jenjang_pendidikan"`
	Nik         string    `db:"nik" json:"nik"`
	TanggalLahir string `db:"tanggal_lahir" json:"tanggal_lahir"`
	TempatLahir string    `db:"tempat_lahir" json:"tempat_lahir"`
}

type Pendaftaran struct {
	Id          int64     `db:"id" json:"id"`
	IdBeasiswa  int64     `db:"id_beasiswa" json:"id_beasiswa"`
	IdSiswa     int64     `db:"id_siswa" json:"id_siswa"`
	TanggalDaftar string `db:"tanggal_daftar" json:"tanggal_daftar"`
	Status      string    `db:"status" json:"status"`
}

type Mitra struct {
	Id          int64     `db:"id" json:"id"`
	Nama        string    `db:"nama" json:"nama"`
	Email       string    `db:"email" json:"email"`
	Lokasi      string    `db:"lokasi" json:"lokasi"`
	NoTelp      string    `db:"no_telp" json:"no_telp"`
	Legalitas   string    `db:"legalitas" json:"legalitas"`
}

