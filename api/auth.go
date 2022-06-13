package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type Siswa struct {
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	Nama              string `json:"nama" validate:"required"`
	JenjangPendidikan string `json:"jenjang_pendidikan" validate:"required"`
	Nik               string `json:"nik" validate:"required"`
	TempatLahir       string `json:"tempat_lahir" validate:"required"`
	TanggalLahir      string `json:"tanggal_lahir" validate:"required"`
}

type LoginSiswa struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginSuccessResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

var jwtKey = []byte("key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (api *API) login(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var s LoginSiswa
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	res, err := api.siswaRepo.Login(s.Email, s.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	expTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Email: res.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expTime,
	})

	json.NewEncoder(w).Encode(LoginSuccessResponse{Email: res.Email, Token: tokenString})
}

func (api *API) register(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	var s Siswa
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validator.New().Struct(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	res, err := api.siswaRepo.Register(s.Nama, s.Password, s.Email, s.JenjangPendidikan, s.Nik, s.TempatLahir, s.TanggalLahir)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (api *API) logout(w http.ResponseWriter, r *http.Request) {
	api.AllowOrigin(w, r)
	token, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logout success"))
}
