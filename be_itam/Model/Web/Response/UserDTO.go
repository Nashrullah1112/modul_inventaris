package Response

import (
	"itam/Model/Domain"
	"time"
)

// UserCreateRequest represents the request payload for creating a new User
type UserCreateRequest struct {
	NIP       int64  `json:"nip" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	JabatanID int64  `json:"jabatan_id" validate:"required"`
	DivisiID  int64  `json:"divisi_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

// UserUpdateRequest represents the request payload for updating an existing User
type UserUpdateRequest struct {
	ID        int64  `json:"id" validate:"required"`
	NIP       int64  `json:"nip" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	JabatanID int64  `json:"jabatan_id" validate:"required"`
	DivisiID  int64  `json:"divisi_id" validate:"required"`
}

// UserResponse represents the response payload for returning User data
type UserResponse struct {
	ID               int64           `json:"id"`
	NIP              int64           `json:"nip"`
	Email            string          `json:"email"`
	JabatanID        int64           `json:"jabatan_id"`
	Jabatan          JabatanResponse `json:"jabatan"`
	DivisiID         int64           `json:"divisi_id"`
	Divisi           DivisiResponse  `json:"divisi"`
	TanggalBergabung time.Time       `json:"tanggal_bergabung"`
	Nama             string          `json:"nama"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token Domain.JwtTokenDetail `json:"token"`
}
