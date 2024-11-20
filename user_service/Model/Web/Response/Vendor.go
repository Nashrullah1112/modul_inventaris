package Response

// VendorCreateRequest represents the request payload for creating a new Vendor
type VendorCreateRequest struct {
	Nama string `json:"nama" validate:"required"`
}

// VendorUpdateRequest represents the request payload for updating an existing Vendor
type VendorUpdateRequest struct {
	ID   int64  `json:"id" validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

// VendorResponse represents the response payload for returning Vendor data
type VendorResponse struct {
	ID   int64  `json:"id"`
	Nama string `json:"nama"`
}
