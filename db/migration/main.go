package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./beasiswa.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS beasiswa (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nama TEXT,
		jenis_beasiswa TEXT,
		jenjang_pendidikan TEXT,
		tanggal_mulai TEXT,
		tanggal_selesai TEXT);
		
		CREATE TABLE IF NOT EXISTS siswa (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nama TEXT,
			password TEXT,
			email TEXT,
			jenjang_pendidikan TEXT,
			nik TEXT,
			tanggal_lahir TEXT,
			tempat_lahir TEXT);
		
		CREATE TABLE IF NOT EXISTS pendaftaran (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			id_beasiswa INTEGER,
			id_siswa INTEGER,
			tanggal_daftar TEXT,
			status TEXT);
		
		CREATE TABLE IF NOT EXISTS mitra (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nama TEXT,
			email TEXT,
			lokasi TEXT,
			no_telp TEXT,
			legalitas TEXT);
			
		INSERT INTO beasiswa (nama, jenis_beasiswa, jenjang_pendidikan, tanggal_mulai, tanggal_selesai)
		VALUES ('Beasiswa Pertama', 'Dalam Negeri', 'S1', '2020-01-01', '2020-01-01'),
		('Beasiswa Kedua', 'Luar Negeri', 'S1', '2020-01-01', '2020-01-01');
		
		INSERT INTO siswa (nama, password, email, jenjang_pendidikan, nik, tanggal_lahir, tempat_lahir)
		VALUES ('Siswa Pertama', '12345', 'ex@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta'),
		('Siswa Kedua', '12345', 'contoh@gmail.com', 'S1', '123456789', '2020-01-01', 'Jakarta');`)

	if err != nil {
		panic(err)
	}
}
