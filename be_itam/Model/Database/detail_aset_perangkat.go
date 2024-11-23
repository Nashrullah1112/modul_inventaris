package Database

import (
	"time"
)

const TableNameDetailAsetPerangkat = "detail_aset_perangkat"

// DetailAsetPerangkat mapped from table <detail_aset_perangkat>
type DetailAsetPerangkat struct {
	ID                   int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	LokasiPenerima       string    `gorm:"column:lokasi_penerima;not null" json:"lokasi_penerima"`
	WaktuPenerimaan      time.Time `gorm:"column:waktu_penerimaan;not null" json:"waktu_penerimaan"`
	TandaTerima          string    `gorm:"column:tanda_terima;not null" json:"tanda_terima"`
	TipeAset             string    `gorm:"column:tipe_aset;not null" json:"tipe_aset"`
	WaktuAktivasiAset    time.Time `gorm:"column:waktu_aktivasi_aset;not null" json:"waktu_aktivasi_aset"`
	HasilPemeriksaanAset string    `gorm:"column:hasil_pemeriksaan_aset;not null" json:"hasil_pemeriksaan_aset"`
	SerialNumber         string    `gorm:"column:serial_number;not null" json:"serial_number"`
	Model                string    `gorm:"column:model;not null" json:"model"`
	MasaGaransiMulai     time.Time `gorm:"column:masa_garansi_mulai;not null" json:"masa_garansi_mulai"`
	NomorKartuGaransi    string    `gorm:"column:nomor_kartu_garansi;not null" json:"nomor_kartu_garansi"`
	Prosesor             string    `gorm:"column:prosesor" json:"prosesor"`
	KapasitasRAM         int32     `gorm:"column:kapasitas_ram" json:"kapasitas_ram"`
	KapasitasRom         int32     `gorm:"column:kapasitas_rom" json:"kapasitas_rom"`
	TipeRAM              string    `gorm:"column:tipe_ram" json:"tipe_ram"`
	TipePenyimpnanan     string    `gorm:"column:tipe_penyimpnanan;not null" json:"tipe_penyimpnanan"`
	StatusAset           string    `gorm:"column:status_aset;not null" json:"status_aset"`
	NilaiAset            int64     `gorm:"column:nilai_aset;not null" json:"nilai_aset"`
	NilaiDepresiasi      int64     `gorm:"column:nilai_depresiasi;not null" json:"nilai_depresiasi"`
	JangkaMasaPakai      int32     `gorm:"column:jangka_masa_pakai;not null" json:"jangka_masa_pakai"`
	WaktuAsetKeluar      time.Time `gorm:"column:waktu_aset_keluar;not null" json:"waktu_aset_keluar"`
	KondisiAsetKeluar    string    `gorm:"column:kondisi_aset_keluar;not null" json:"kondisi_aset_keluar"`
	NotaPembelian        string    `gorm:"column:nota_pembelian;not null" json:"nota_pembelian"`
	DivisiID             int64     `gorm:"column:divisi_id;not null" json:"divisi_id"`
	Divisi               Divisi    `gorm:"foreignKey:DivisiID;references:ID" json:"divisi"`
	UserID               int64     `gorm:"column:user_id;not null" json:"user_id"`
	User                 User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	AssetID              int64     `gorm:"column:asset_id;not null" json:"asset_id"`
	Asset                Asset     `gorm:"foreignKey:AssetID;references:ID" json:"asset"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName DetailAsetPerangkat's table name
func (*DetailAsetPerangkat) TableName() string {
	return TableNameDetailAsetPerangkat
}
