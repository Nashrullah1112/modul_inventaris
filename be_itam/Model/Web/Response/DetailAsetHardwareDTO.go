package Response

import "time"

// DetailAsetHardwareCreateRequest digunakan untuk membuat data baru DetailAsetHardware
type DetailAsetHardwareCreateRequest struct {
	WaktuPenerimaan          string `json:"waktu_penerimaan" validate:"required"`
	BuktiPenerimaan          string `json:"bukti_penerimaan" validate:"required"`
	TipeAset                 string `json:"tipe_aset" validate:"required"`
	TanggalAktivasiPerangkat string `json:"tanggal_aktivasi_perangkat" validate:"required"`
	HasilPemeriksaan         string `json:"hasil_pemeriksaan" validate:"required"`
	SerialNumber             string `json:"serial_number" validate:"required"`
	Model                    string `json:"model" validate:"required"`
	WaktuGaransiMulai        string `json:"waktu_garansi_mulai" validate:"required"`
	WaktuGaransiBerakhir     string `json:"waktu_garansi_berakhir" validate:"required"`
	NomorKartuGaransi        string `json:"nomor_kartu_garansi" validate:"required"`
	SpesifikasiPerangkat     string `json:"spesifikasi_perangkat" validate:"required"`
	StatusAset               string `json:"status_aset" validate:"required"`
	PenanggungjawabAset      string `json:"penanggungjawab_aset" validate:"required"`
	LokasiPenyimpananID      string `json:"lokasi_penyimpanan_id" validate:"required"`
	JangkaMasaPakai          int32  `json:"jangka_masa_pakai" validate:"required"`
	WaktuAsetKeluar          string `json:"waktu_aset_keluar" validate:"required"`
	KondisiAset              string `json:"kondisi_aset" validate:"required"`
	NotaPembelian            string `json:"nota_pembelian" validate:"required"`
	DivisiID                 int64  `json:"divisi_id" validate:"required"`
	AssetID                  int64  `json:"asset_id" validate:"required"`
	VendorID                 int64  `json:"vendor_id" validate:"required"`
}

// DetailAsetHardwareUpdateRequest digunakan untuk memperbarui data DetailAsetHardware
type DetailAsetHardwareUpdateRequest struct {
	ID                       int64  `json:"id" validate:"required"`
	WaktuPenerimaan          string `json:"waktu_penerimaan" validate:"required"`
	BuktiPenerimaan          string `json:"bukti_penerimaan" validate:"required"`
	TipeAset                 string `json:"tipe_aset" validate:"required"`
	TanggalAktivasiPerangkat string `json:"tanggal_aktivasi_perangkat" validate:"required"`
	HasilPemeriksaan         string `json:"hasil_pemeriksaan" validate:"required"`
	SerialNumber             string `json:"serial_number" validate:"required"`
	Model                    string `json:"model" validate:"required"`
	WaktuGaransiMulai        string `json:"waktu_garansi_mulai" validate:"required"`
	WaktuGaransiBerakhir     string `json:"waktu_garansi_berakhir" validate:"required"`
	NomorKartuGaransi        string `json:"nomor_kartu_garansi" validate:"required"`
	SpesifikasiPerangkat     string `json:"spesifikasi_perangkat" validate:"required"`
	StatusAset               string `json:"status_aset" validate:"required"`
	PenanggungjawabAset      string `json:"penanggungjawab_aset" validate:"required"`
	LokasiPenyimpananID      string `json:"lokasi_penyimpanan_id" validate:"required"`
	JangkaMasaPakai          int32  `json:"jangka_masa_pakai" validate:"required"`
	WaktuAsetKeluar          string `json:"waktu_aset_keluar" validate:"required"`
	KondisiAset              string `json:"kondisi_aset" validate:"required"`
	NotaPembelian            string `json:"nota_pembelian" validate:"required"`
	DivisiID                 int64  `json:"divisi_id" validate:"required"`
	AssetID                  int64  `json:"asset_id" validate:"required"`
	VendorID                 int64  `json:"vendor_id" validate:"required"`
}

// DetailAsetHardwareResponse digunakan untuk menampilkan data DetailAsetHardware
type DetailAsetHardwareResponse struct {
	Id                       int64         `json:"id"`
	WaktuPenerimaan          string        `json:"waktu_penerimaan"`
	BuktiPenerimaan          string        `json:"bukti_penerimaan"`
	TipeAset                 string        `json:"tipe_aset"`
	TanggalAktivasiPerangkat string        `json:"tanggal_aktivasi_perangkat"`
	HasilPemeriksaan         string        `json:"hasil_pemeriksaan"`
	SerialNumber             string        `json:"serial_number"`
	Model                    string        `json:"model"`
	WaktuGaransiMulai        string        `json:"waktu_garansi_mulai"`
	WaktuGaransiBerakhir     string        `json:"waktu_garansi_berakhir"`
	NomorKartuGaransi        string        `json:"nomor_kartu_garansi"`
	SpesifikasiPerangkat     string        `json:"spesifikasi_perangkat"`
	StatusAset               string        `json:"status_aset"`
	PenanggungjawabAset      string        `json:"penanggungjawab_aset"`
	LokasiPenyimpananID      string        `json:"lokasi_penyimpanan_id"`
	JangkaMasaPakai          int32         `json:"jangka_masa_pakai"`
	WaktuAsetKeluar          time.Time     `json:"waktu_aset_keluar"`
	KondisiAset              string        `json:"kondisi_aset"`
	NotaPembelian            string        `json:"nota_pembelian"`
	DivisiID                 int64         `json:"divisi_id"`
	AssetID                  int64         `json:"asset_id"`
	Asset                    AssetResponse `json:"asset"`
}
type AssetHardwareCreateRequest struct {
	MerekPerangkat            string  `form:"merek_perangkat"`
	VendorID                  int64   `form:"vendor_id"`
	NomorNota                 string  `form:"nomor_nota"`
	LokasiPenerima            string  `form:"lokasi_penerima"`
	TanggalPenerimaan         string  `form:"tanggal_penerimaan"` // Format ISO-8601 (2024-10-12T10:30:00)
	MasaGaransiMulai          string  `form:"masa_garansi_mulai"` // Format ISO-8601 (2024-10-12)
	PenanggungJawabPerangkat  string  `form:"penanggung_jawab_perangkat"`
	Model                     string  `form:"model"`
	SerialNumber              string  `form:"serial_number"`
	HargaPerangkat            float64 `form:"harga_perangkat"`
	KondisiAset               string  `form:"kondisi_aset"`
	TipePerangkat             string  `form:"tipe_perangkat"`
	TanggalAktivasiPerangkat  string  `form:"tanggal_aktivasi_perangkat"` // Format ISO-8601 (2024-10-13)
	MasaBerakhirGaransi       string  `form:"masa_berakhir_garansi"`      // Format ISO-8601 (2025-10-12)
	HasilPemeriksaanPerangkat string  `form:"hasil_pemeriksaan_perangkat"`
	JangkaMasaPakai           int32   `form:"jangka_masa_pakai"`
	StatusPerangkat           string  `form:"status_perangkat"`
	TanggalAsetKeluar         string  `form:"tanggal_aset_keluar"` // Format ISO-8601 (2024-10-13)
	DivisiPengguna            int64   `form:"divisi_pengguna"`
	DetailSpesifikasi         string  `form:"detail_spesifikasi"`
	NomorKartuGaransi         string  `form:"nomor_kartu_garansi"`
}
type AssetHardwareUpdateRequest struct {
	MerekPerangkat            string  `form:"merek_perangkat"`
	VendorID                  int64   `form:"vendor_id"`
	NomorNota                 string  `form:"nomor_nota"`
	LokasiPenerima            string  `form:"lokasi_penerima"`
	TanggalPenerimaan         string  `form:"tanggal_penerimaan"`
	MasaGaransiMulai          string  `form:"masa_garansi_mulai"`
	PenanggungJawabPerangkat  string  `form:"penanggung_jawab_perangkat"`
	Model                     string  `form:"model"`
	SerialNumber              string  `form:"serial_number"`
	HargaPerangkat            float64 `form:"harga_perangkat"`
	KondisiAset               string  `form:"kondisi_aset"`
	TipePerangkat             string  `form:"tipe_perangkat"`
	TanggalAktivasiPerangkat  string  `form:"tanggal_aktivasi_perangkat"`
	MasaBerakhirGaransi       string  `form:"masa_berakhir_garansi"`
	HasilPemeriksaanPerangkat string  `form:"hasil_pemeriksaan_perangkat"`
	JangkaMasaPakai           int32   `form:"jangka_masa_pakai"`
	StatusPerangkat           string  `form:"status_perangkat"`
	TanggalAsetKeluar         string  `form:"tanggal_aset_keluar"`
	DivisiPengguna            int64   `form:"divisi_pengguna"`
	DetailSpesifikasi         string  `form:"detail_spesifikasi"`
	NomorKartuGaransi         string  `form:"nomor_kartu_garansi"`
}
