package Database

import (
	"time"
)

const TableNameDetailAsetHardware = "detail_aset_hardware"

// DetailAsetHardware mapped from table <detail_aset_hardware>
type DetailAsetHardware struct {
	ID                       int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	WaktuPenerimaan          string    `gorm:"column:waktu_penerimaan;not null" json:"waktu_penerimaan"`
	BuktiPenerimaan          string    `gorm:"column:bukti_penerimaan;not null" json:"bukti_penerimaan"`
	TipeAset                 string    `gorm:"column:tipe_aset;not null" json:"tipe_aset"`
	TanggalAktivasiPerangkat string    `gorm:"column:tanggal_aktivasi_perangkat;not null" json:"tanggal_aktivasi_perangkat"`
	HasilPemeriksaan         string    `gorm:"column:hasil_pemeriksaan;not null" json:"hasil_pemeriksaan"`
	SerialNumber             string    `gorm:"column:serial_number;not null;unique" json:"serial_number"`
	Model                    string    `gorm:"column:model;not null" json:"model"`
	WaktuGaransiMulai        string    `gorm:"column:waktu_garansi_mulai;not null" json:"waktu_garansi_mulai"`
	WaktuGaransiBerakhir     string    `gorm:"column:waktu_garansi_berakhir;not null" json:"waktu_garansi_berakhir"`
	NomorKartuGaransi        string    `gorm:"column:nomor_kartu_garansi;not null" json:"nomor_kartu_garansi"`
	SpesifikasiPerangkat     string    `gorm:"column:spesifikasi_perangkat;not null" json:"spesifikasi_perangkat"`
	StatusAset               string    `gorm:"column:status_aset;not null" json:"status_aset"`
	PenanggungjawabAset      string    `gorm:"column:penanggungjawab_aset;not null" json:"penanggungjawab_aset"`
	LokasiPenyimpananID      string    `gorm:"column:lokasi_penyimpanan_id;not null" json:"lokasi_penyimpanan_id"`
	JangkaMasaPakai          int32     `gorm:"column:jangka_masa_pakai;not null" json:"jangka_masa_pakai"`
	WaktuAsetKeluar          time.Time `gorm:"column:waktu_aset_keluar;not null" json:"waktu_aset_keluar"`
	KondisiAset              string    `gorm:"column:kondisi_aset;not null" json:"kondisi_aset"`
	NotaPembelian            string    `gorm:"column:nota_pembelian;not null" json:"nota_pembelian"`
	DivisiID                 int64     `gorm:"column:divisi_id;not null" json:"divisi_id"`
	Divisi                   Divisi    `gorm:"foreignKey:DivisiID;references:ID" json:"divisi"`
	AssetID                  int64     `gorm:"column:asset_id;not null" json:"asset_id"`
	Asset                    Asset     `gorm:"foreignKey:AssetID;references:ID" json:"asset"`
	CreatedAt                time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt                time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName DetailAsetHardware's table name
func (*DetailAsetHardware) TableName() string {
	return TableNameDetailAsetHardware
}
