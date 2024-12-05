package Response

type AssetResponse struct {
	Id           int64          `json:"id"`
	VendorID     int64          `json:"vendor_id"`
	SerialNumber string         `json:"serial_number"`
	Merk         string         `json:"merk"`
	Model        string         `json:"model"`
	NomorNota    string         `json:"nomor_nota"`
	Status       *string        `json:"status,omitempty"`
	Vendor       VendorResponse `json:"vendor"`
}

type AssetCreateRequest struct {
	VendorID     int64  `json:"vendor_id"`
	SerialNumber string `json:"serial_number"`
	Merk         string `json:"merk"`
	Model        string `json:"model"`
	NomorNota    string `json:"nomor_nota"`
}

type AssetUpdateRequest struct {
	Id           int64  `json:"id"`
	VendorID     int64  `json:"vendor_id"`
	SerialNumber string `json:"serial_number"`
	Merk         string `json:"merk"`
	Model        string `json:"model"`
	NomorNota    string `json:"nomor_nota"`
}

type IDAsset struct {
	Id int64 `json:"id"`
}
