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
	Id     int64  `json:"id"`
	Status string `json:"status"`
}
type DetailAssetResponse struct {
	Id           int64                  `json:"id"`
	VendorID     int64                  `json:"vendor_id"`
	SerialNumber string                 `json:"serial_number,omitempty"`
	Merk         string                 `json:"merk,omitempty"`
	Model        string                 `json:"model,omitempty"`
	NomorNota    string                 `json:"nomor_nota,omitempty"`
	Status       string                 `json:"status,omitempty"`
	Vendor       *VendorResponse        `json:"vendor,omitempty"`
	Hardware     *AsetHardwareResponse  `json:"hardware,omitempty"`
	Lisensi      *AsetLisensiResponse   `json:"software,omitempty"`
	Perangkat    *AsetPerangkatResponse `json:"perangkat,omitempty"`
	Aplikasi     *AsetAplikasiResponse  `json:"aplikasi,omitempty"`
}
