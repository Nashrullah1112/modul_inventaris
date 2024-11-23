package Response

// JabatanCreateRequest represents the request payload for creating a new Jabatan
type JabatanCreateRequest struct {
	Nama string `json:"nama" validate:"required"`
}

// JabatanUpdateRequest represents the request payload for updating an existing Jabatan
type JabatanUpdateRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

// JabatanResponse represents the response payload for returning Jabatan data
type JabatanResponse struct {
	ID   int64  `json:"id"`
	Nama string `json:"nama"`
}
