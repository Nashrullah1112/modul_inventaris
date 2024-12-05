package Response

import (
	"time"
)

type (
	DetailAsetPerangkatCreateRequest struct {
		VendorID          int64     `form:"vendor_id"`
		SerialNumber      string    `form:"serial_number"`
		Merk              string    `form:"merk"`
		Model             string    `form:"model"`
		NomorNota         string    `form:"nomor_nota"`
		LokasiPenerima    string    `form:"lokasi_penerima"`
		WaktuPenerimaan   time.Time `form:"waktu_penerimaan"`
		TipeAset          string    `form:"tipe_aset"`
		WaktuAktivasiAset time.Time `form:"waktu_aktivasi_aset"`
		MasaGaransiMulai  time.Time `form:"masa_garansi_mulai"`
		NomorKartuGaransi string    `form:"nomor_kartu_garansi"`
		HasilPemeriksaan  string    `form:"hasil_pemeriksaan"`
		Prosesor          string    `form:"prosesor"`
		KapasitasRAM      int32     `form:"kapasitas_ram"`
		KapasitasRom      int32     `form:"kapasitas_rom"`
		TipeRAM           string    `form:"tipe_ram"`
		TipePenyimpanan   string    `form:"tipe_penyimpnanan"`
		StatusAset        string    `form:"status_aset"`
		NilaiAset         int64     `form:"nilai_aset"`
		NilaiSisa         int64     `form:"nilai_sisa"`
		JangkaMasaPakai   int32     `form:"jangka_masa_pakai"`
		WaktuAsetKeluar   time.Time `form:"waktu_aset_keluar"`
		KondisiAsetKeluar string    `form:"kondisi_aset_keluar"`
		DivisiID          int64     `form:"divisi_pengguna"`
		UserID            int64     `form:"user_id"`
	}

	DetailAsetPerangkatUpdateRequest struct {
		LokasiPenerima       string    `json:"lokasi_penerima" validate:"required"`
		WaktuPenerimaan      time.Time `json:"waktu_penerimaan" validate:"required"`
		TandaTerima          string    `json:"tanda_terima" validate:"required"`
		TipeAset             string    `json:"tipe_aset" validate:"required"`
		WaktuAktivasiAset    time.Time `json:"waktu_aktivasi_aset" validate:"required"`
		HasilPemeriksaanAset string    `json:"hasil_pemeriksaan_aset" validate:"required"`
		SerialNumber         string    `json:"serial_number" validate:"required"`
		Model                string    `json:"model" validate:"required"`
		MasaGaransiMulai     time.Time `json:"masa_garansi_mulai" validate:"required"`
		NomorKartuGaransi    string    `json:"nomor_kartu_garansi" validate:"required"`
		Prosesor             string    `json:"prosesor" validate:"required"`
		KapasitasRAM         int32     `json:"kapasitas_ram" validate:"required"`
		KapasitasRom         int32     `json:"kapasitas_rom" validate:"required"`
		TipeRAM              string    `json:"tipe_ram" validate:"required"`
		TipePenyimpanan      string    `json:"tipe_penyimpnanan" validate:"required"`
		StatusAset           string    `json:"status_aset" validate:"required"`
		NilaiAset            int64     `json:"nilai_aset" validate:"required"`
		NilaiDepresiasi      int64     `json:"nilai_depresiasi" validate:"required"`
		JangkaMasaPakai      int32     `json:"jangka_masa_pakai" validate:"required"`
		WaktuAsetKeluar      time.Time `json:"waktu_aset_keluar" validate:"required"`
		KondisiAsetKeluar    string    `json:"kondisi_aset_keluar" validate:"required"`
		NotaPembelian        string    `json:"nota_pembelian" validate:"required"`
		DivisiID             int64     `json:"divisi_id" validate:"required"`
		UserID               int64     `json:"user_id" validate:"required"`
		VendorID             int64     `json:"vendor_id" validate:"required"`
		Merk                 string    `json:"merk" validate:"required"`
		NomorNota            string    `json:"nomor_nota" validate:"required"`
	}

	DetailAsetPerangkatResponse struct {
		ID                   int64         `json:"id" `
		LokasiPenerima       string        `json:"lokasi_penerima" `
		WaktuPenerimaan      time.Time     `json:"waktu_penerimaan" `
		TandaTerima          string        `json:"tanda_terima" `
		TipeAset             string        `json:"tipe_aset"`
		WaktuAktivasiAset    time.Time     `json:"waktu_aktivasi_aset"`
		HasilPemeriksaanAset string        `json:"hasil_pemeriksaan_aset"`
		SerialNumber         string        `json:"serial_number"`
		Model                string        `json:"model"`
		MasaGaransiMulai     time.Time     `json:"masa_garansi_mulai"`
		NomorKartuGaransi    string        `json:"nomor_kartu_garansi"`
		Prosesor             string        `json:"prosesor"`
		KapasitasRAM         int32         `json:"kapasitas_ram"`
		KapasitasRom         int32         `json:"kapasitas_rom"`
		TipeRAM              string        `json:"tipe_ram"`
		TipePenyimpanan      string        `json:"tipe_penyimpnanan"`
		StatusAset           string        `json:"status_aset"`
		NilaiAset            int64         `json:"nilai_aset"`
		NilaiDepresiasi      int64         `json:"nilai_depresiasi"`
		JangkaMasaPakai      int32         `json:"jangka_masa_pakai"`
		WaktuAsetKeluar      time.Time     `json:"waktu_aset_keluar"`
		KondisiAsetKeluar    string        `json:"kondisi_aset_keluar"`
		NotaPembelian        string        `json:"nota_pembelian"`
		DivisiID             int64         `json:"divisi_id"`
		UserID               int64         `json:"user_id"`
		AssetID              int64         `json:"asset_id"`
		Asset                AssetResponse `json:"asset"`
	}
	AsetPerangkatResponse struct {
		ID                   int64     `json:"id" `
		LokasiPenerima       string    `json:"lokasi_penerima" `
		WaktuPenerimaan      time.Time `json:"waktu_penerimaan" `
		TandaTerima          string    `json:"tanda_terima" `
		TipeAset             string    `json:"tipe_aset"`
		WaktuAktivasiAset    time.Time `json:"waktu_aktivasi_aset"`
		HasilPemeriksaanAset string    `json:"hasil_pemeriksaan_aset"`
		SerialNumber         string    `json:"serial_number"`
		Model                string    `json:"model"`
		MasaGaransiMulai     time.Time `json:"masa_garansi_mulai"`
		NomorKartuGaransi    string    `json:"nomor_kartu_garansi"`
		Prosesor             string    `json:"prosesor"`
		KapasitasRAM         int32     `json:"kapasitas_ram"`
		KapasitasRom         int32     `json:"kapasitas_rom"`
		TipeRAM              string    `json:"tipe_ram"`
		TipePenyimpanan      string    `json:"tipe_penyimpnanan"`
		StatusAset           string    `json:"status_aset"`
		NilaiAset            int64     `json:"nilai_aset"`
		NilaiDepresiasi      int64     `json:"nilai_depresiasi"`
		JangkaMasaPakai      int32     `json:"jangka_masa_pakai"`
		WaktuAsetKeluar      time.Time `json:"waktu_aset_keluar"`
		KondisiAsetKeluar    string    `json:"kondisi_aset_keluar"`
		NotaPembelian        string    `json:"nota_pembelian"`
		DivisiID             int64     `json:"divisi_id"`
		UserID               int64     `json:"user_id"`
		AssetID              int64     `json:"asset_id"`
	}
)
