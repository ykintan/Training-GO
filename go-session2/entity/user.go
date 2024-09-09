package entity

import "time"

type User struct {
	ID        int       `json:"id"`         // ID pengguna sebagai primary key
	Name      string    `json:"name"`       // Nama pengguna (wajib diisi)
	Email     string    `json:"email"`      // Email pengguna (wajib diisi, harus unik)
	Password  string    `json:"password"`   // Kata sandi pengguna (wajib diisi)
	CreatedAt time.Time `json:"created_at"` // Waktu pembuatan pengguna
	UpdatedAt time.Time `json:"updated_at"`
}
