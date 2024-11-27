package Response

import (
	"time"
)

type DetaiAsetAplikasiResponse struct {
	Id                      int64         `json:"id"`
	NamaAplikasi            string        `json:"nama_aplikasi"`
	TanggalPembuatan        time.Time     `json:"tanggal_pembuatan"`
	TanggalTerima           time.Time     `json:"tanggal_terima"`
	LokasiServerPenyimpanan string        `json:"lokasi_server_penyimpanan"`
	TipeAplikasi            string        `json:"tipe_aplikasi"`
	LinkAplikasi            string        `json:"link_aplikasi"`
	SertifikasiAplikasi     string        `json:"sertifikasi_aplikasi,omitempty"`
	TanggalAktif            time.Time     `json:"tanggal_aktif"`
	TanggalKadaluarsa       time.Time     `json:"tanggal_kadaluarsa"`
	AssetID                 int64         `json:"asset_id"`
	Asset                   AssetResponse `json:"asset"`
}

type DetaiAsetAplikasiCreateRequest struct {
	VendorID                int64     `json:"vendor_id"`
	SerialNumber            string    `json:"serial_number"`
	Merk                    string    `json:"merk"`
	Model                   string    `json:"model"`
	NomorNota               string    `json:"nomor_nota"`
	NamaAplikasi            string    `json:"nama_aplikasi" validate:"required"`
	TanggalPembuatan        time.Time `json:"tanggal_pembuatan" validate:"required"`
	TanggalTerima           time.Time `json:"tanggal_terima" validate:"required"`
	LokasiServerPenyimpanan string    `json:"lokasi_server_penyimpanan" validate:"required"`
	TipeAplikasi            string    `json:"tipe_aplikasi" validate:"required"`
	LinkAplikasi            string    `json:"link_aplikasi" validate:"required"`
	SertifikasiAplikasi     string    `json:"sertifikasi_aplikasi,omitempty"`
	TanggalAktif            time.Time `json:"tanggal_aktif" validate:"required"`
	TanggalKadaluarsa       time.Time `json:"tanggal_kadaluarsa" validate:"required"`
}

type DetaiAsetAplikasiUpdateRequest struct {
	NamaAplikasi            string    `json:"nama_aplikasi" validate:"required"`
	TanggalPembuatan        time.Time `json:"tanggal_pembuatan" validate:"required"`
	TanggalTerima           time.Time `json:"tanggal_terima" validate:"required"`
	LokasiServerPenyimpanan string    `json:"lokasi_server_penyimpanan" validate:"required"`
	TipeAplikasi            string    `json:"tipe_aplikasi" validate:"required"`
	LinkAplikasi            string    `json:"link_aplikasi" validate:"required"`
	SertifikasiAplikasi     string    `json:"sertifikasi_aplikasi,omitempty"`
	TanggalAktif            time.Time `json:"tanggal_aktif" validate:"required"`
	TanggalKadaluarsa       time.Time `json:"tanggal_kadaluarsa" validate:"required"`
	VendorID                int64     `json:"vendor_id" validate:"required"`
	SerialNumber            string    `json:"serial_number"`
	Merk                    string    `json:"merk"`
	Model                   string    `json:"model"`
	NomorNota               string    `json:"nomor_nota"`
}
