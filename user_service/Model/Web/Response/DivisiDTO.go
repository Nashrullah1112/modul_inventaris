package Response

// DivisiCreateRequest is the DTO for creating a new Divisi.
type DivisiCreateRequest struct {
	Nama string `json:"nama" validate:"required"`
}

// DivisiUpdateRequest is the DTO for updating an existing Divisi.
type DivisiUpdateRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

// DivisiResponse is the DTO for returning Divisi details.
type DivisiResponse struct {
	ID   int64  `json:"id"`
	Nama string `json:"nama"`
}
