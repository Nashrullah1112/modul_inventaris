package Response

import (
	"time"
)

type (
	DetailAsetPerangkatCreateRequest struct {
		VendorID             int64     `json:"vendor_id"`
		SerialNumber         string    `json:"serial_number"`
		Merk                 string    `json:"merk"`
		Model                string    `json:"model"`
		NomorNota            string    `json:"nomor_nota"`
		LokasiPenerima       string    `json:"lokasi_penerima" binding:"required"`
		WaktuPenerimaan      time.Time `json:"waktu_penerimaan" binding:"required"`
		TandaTerima          string    `json:"tanda_terima" binding:"required"`
		TipeAset             string    `json:"tipe_aset" binding:"required"`
		WaktuAktivasiAset    time.Time `json:"waktu_aktivasi_aset" binding:"required"`
		HasilPemeriksaanAset string    `json:"hasil_pemeriksaan_aset" binding:"required"`
		MasaGaransiMulai     time.Time `json:"masa_garansi_mulai" binding:"required"`
		NomorKartuGaransi    string    `json:"nomor_kartu_garansi" binding:"required"`
		Prosesor             string    `json:"prosesor"`
		KapasitasRAM         int32     `json:"kapasitas_ram"`
		KapasitasRom         int32     `json:"kapasitas_rom"`
		TipeRAM              string    `json:"tipe_ram"`
		TipePenyimpnanan     string    `json:"tipe_penyimpnanan" binding:"required"`
		StatusAset           string    `json:"status_aset" binding:"required"`
		NilaiAset            int64     `json:"nilai_aset" binding:"required"`
		NilaiSisa            int64     `json:"nilai_sisa" binding:"required"`
		JangkaMasaPakai      int32     `json:"jangka_masa_pakai" binding:"required"`
		WaktuAsetKeluar      time.Time `json:"waktu_aset_keluar" binding:"required"`
		KondisiAsetKeluar    string    `json:"kondisi_aset_keluar" binding:"required"`
		NotaPembelian        string    `json:"nota_pembelian" binding:"required"`
		DivisiID             int64     `json:"divisi_id" binding:"required"`
		UserID               int64     `json:"user_id" binding:"required"`
	}

	DetailAsetPerangkatUpdateRequest struct {
		LokasiPenerima       string    `json:"lokasi_penerima" binding:"required"`
		WaktuPenerimaan      time.Time `json:"waktu_penerimaan" binding:"required"`
		TandaTerima          string    `json:"tanda_terima" binding:"required"`
		TipeAset             string    `json:"tipe_aset" binding:"required"`
		WaktuAktivasiAset    time.Time `json:"waktu_aktivasi_aset" binding:"required"`
		HasilPemeriksaanAset string    `json:"hasil_pemeriksaan_aset" binding:"required"`
		SerialNumber         string    `json:"serial_number" binding:"required"`
		Model                string    `json:"model" binding:"required"`
		MasaGaransiMulai     time.Time `json:"masa_garansi_mulai" binding:"required"`
		NomorKartuGaransi    string    `json:"nomor_kartu_garansi" binding:"required"`
		Prosesor             string    `json:"prosesor"`
		KapasitasRAM         int32     `json:"kapasitas_ram"`
		KapasitasRom         int32     `json:"kapasitas_rom"`
		TipeRAM              string    `json:"tipe_ram"`
		TipePenyimpnanan     string    `json:"tipe_penyimpnanan" binding:"required"`
		StatusAset           string    `json:"status_aset" binding:"required"`
		NilaiAset            int64     `json:"nilai_aset" binding:"required"`
		NilaiDepresiasi      int64     `json:"nilai_depresiasi" binding:"required"`
		JangkaMasaPakai      int32     `json:"jangka_masa_pakai" binding:"required"`
		WaktuAsetKeluar      time.Time `json:"waktu_aset_keluar" binding:"required"`
		KondisiAsetKeluar    string    `json:"kondisi_aset_keluar" binding:"required"`
		NotaPembelian        string    `json:"nota_pembelian" binding:"required"`
		DivisiID             int64     `json:"divisi_id" binding:"required"`
		UserID               int64     `json:"user_id" binding:"required"`
		VendorID             int64     `json:"vendor_id" binding:"required"`
		Merk                 string    `json:"merk"`
		NomorNota            string    `json:"nomor_nota"`
	}

	DetailAsetPerangkatResponse struct {
		ID                   int64         `json:"id"`
		LokasiPenerima       string        `json:"lokasi_penerima"`
		WaktuPenerimaan      time.Time     `json:"waktu_penerimaan"`
		TandaTerima          string        `json:"tanda_terima"`
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
		TipePenyimpnanan     string        `json:"tipe_penyimpnanan"`
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
)
