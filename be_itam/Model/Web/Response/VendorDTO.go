package Response

// VendorCreateRequest represents the request payload for creating a new Vendor
type VendorCreateRequest struct {
	PIC              string `json:"nama_pic" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	NomorKontak      string `json:"nomor_kontak" validate:"required"`
	LokasiPerusahaan string `json:"lokasi_perusahaan" validate:"required"`
	NomorSIUP        string `json:"nomor_siup" validate:"required"`
	NomorNIB         string `json:"nomor_nib" validate:"required"`
	NomorNPWP        string `json:"nomor_npwp" validate:"required"`
}

// VendorUpdateRequest represents the request payload for updating an existing Vendor
type VendorUpdateRequest struct {
	ID               int64  `json:"id" validate:"required"`
	PIC              string `json:"nama_pic" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	NomorKontak      string `json:"nomor_kontak" validate:"required"`
	LokasiPerusahaan string `json:"lokasi_perusahaan" validate:"required"`
	NomorSIUP        string `json:"nomor_siup" validate:"required"`
	NomorNIB         string `json:"nomor_nib" validate:"required"`
	NomorNPWP        string `json:"nomor_npwp" validate:"required"`
}

// VendorResponse represents the response payload for returning Vendor data
type VendorResponse struct {
	ID               int64  `json:"id"`
	PIC              string `json:"nama_pic" validate:"required"`
	Email            string `json:"email" validate:"required,email"`
	NomorKontak      string `json:"nomor_kontak" validate:"required"`
	LokasiPerusahaan string `json:"lokasi_perusahaan" validate:"required"`
	NomorSIUP        string `json:"nomor_siup" validate:"required"`
	NomorNIB         string `json:"nomor_nib" validate:"required"`
	NomorNPWP        string `json:"nomor_npwp" validate:"required"`
}
