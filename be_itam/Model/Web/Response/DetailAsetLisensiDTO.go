package Response

import (
	"time"
)

// DTO untuk membuat DetailAsetLisensi baru
type DetailAsetLisensiCreateRequest struct {
	WaktuPembelian           string    `json:"waktu_pembelian"`
	SNPerangkatTerpasang     string    `json:"SN_perangkat_terpasang"`
	WaktuAktivasi            string    `json:"waktu_aktivasi"`
	TanggalExpired           time.Time `json:"tanggal_expired"`
	TipeKepemilikanAset      string    `json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string    `json:"kategori_lisensi"`
	VersiLisensi             string    `json:"versi_lisensi"`
	MaksimalUserAplikasi     int32     `json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32     `json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string    `json:"tipe_lisensi"`
}

// DTO untuk mengupdate DetailAsetLisensi yang ada
type DetailAsetLisensiUpdateRequest struct {
	ID                       int64     `json:"id"`
	WaktuPembelian           string    `json:"waktu_pembelian"`
	SNPerangkatTerpasang     string    `json:"SN_perangkat_terpasang"`
	WaktuAktivasi            string    `json:"waktu_aktivasi"`
	TanggalExpired           time.Time `json:"tanggal_expired"`
	TipeKepemilikanAset      string    `json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string    `json:"kategori_lisensi"`
	VersiLisensi             string    `json:"versi_lisensi"`
	MaksimalUserAplikasi     int32     `json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32     `json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string    `json:"tipe_lisensi"`
}

// DTO untuk merespons hasil query atau operasi pada DetailAsetLisensi
type DetailAsetLisensiResponse struct {
	ID                       int64         `json:"id"`
	WaktuPembelian           string        `json:"waktu_pembelian"`
	SNPerangkatTerpasang     string        `json:"SN_perangkat_terpasang"`
	WaktuAktivasi            string        `json:"waktu_aktivasi"`
	TanggalExpired           time.Time     `json:"tanggal_expired"`
	TipeKepemilikanAset      string        `json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string        `json:"kategori_lisensi"`
	VersiLisensi             string        `json:"versi_lisensi"`
	MaksimalUserAplikasi     int32         `json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32         `json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string        `json:"tipe_lisensi"`
	AssetID                  int64         `json:"asset_id"`
	Asset                    AssetResponse `json:"asset"`
}

type AssetLicenseCreateRequest struct {
	VendorID                 int64     `json:"vendor_id"`
	SerialNumber             string    `json:"serial_number"`
	Merk                     string    `json:"merk"`
	Model                    string    `json:"model"`
	NomorNota                string    `json:"nomor_nota"`
	WaktuPembelian           string    `json:"waktu_pembelian"`
	SNPerangkatTerpasang     string    `json:"SN_perangkat_terpasang"`
	WaktuAktivasi            string    `json:"waktu_aktivasi"`
	TanggalExpired           time.Time `json:"tanggal_expired"`
	TipeKepemilikanAset      string    `json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string    `json:"kategori_lisensi"`
	VersiLisensi             string    `json:"versi_lisensi"`
	MaksimalUserAplikasi     int32     `json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32     `json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string    `json:"tipe_lisensi"`
}
type AssetLicenseUpdateRequest struct {
	VendorID                 int64     `json:"vendor_id"`
	SerialNumber             string    `json:"serial_number"`
	Merk                     string    `json:"merk"`
	Model                    string    `json:"model"`
	NomorNota                string    `json:"nomor_nota"`
	WaktuPembelian           string    `json:"waktu_pembelian"`
	SNPerangkatTerpasang     string    `json:"SN_perangkat_terpasang"`
	WaktuAktivasi            string    `json:"waktu_aktivasi"`
	TanggalExpired           time.Time `json:"tanggal_expired"`
	TipeKepemilikanAset      string    `json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string    `json:"kategori_lisensi"`
	VersiLisensi             string    `json:"versi_lisensi"`
	MaksimalUserAplikasi     int32     `json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32     `json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string    `json:"tipe_lisensi"`
}
type AsetLisensiResponse struct {
	ID                       int64         `json:"id"`
	WaktuPembelian           string        `json:"waktu_pembelian"`
	SNPerangkatTerpasang     string        `json:"SN_perangkat_terpasang"`
	WaktuAktivasi            string        `json:"waktu_aktivasi"`
	TanggalExpired           time.Time     `json:"tanggal_expired"`
	TipeKepemilikanAset      string        `json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string        `json:"kategori_lisensi"`
	VersiLisensi             string        `json:"versi_lisensi"`
	MaksimalUserAplikasi     int32         `json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32         `json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string        `json:"tipe_lisensi"`
	AssetID                  int64         `json:"asset_id"`
	Asset                    AssetResponse `json:"asset"`
}
