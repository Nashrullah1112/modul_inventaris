package Migration

import (
	"time"
)

const TableNameDetaiAsetAplikasi = "detai_aset_aplikasi"

type DetaiAsetAplikasi struct {
	ID                      int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NamaAplikasi            string    `gorm:"column:nama aplikasi;not null" json:"nama aplikasi"`
	TanggalPembuatan        time.Time `gorm:"column:tanggal_pembuatan;not null" json:"tanggal_pembuatan"`
	TanggalTerima           time.Time `gorm:"column:tanggal_terima;not null" json:"tanggal_terima"`
	LokasiServerPenyimpanan string    `gorm:"column:lokasi_server_penyimpanan;not null" json:"lokasi_server_penyimpanan"`
	TipeAplikasi            string    `gorm:"column:tipe_aplikasi;not null" json:"tipe_aplikasi"`
	LinkAplikasi            string    `gorm:"column:link_aplikasi;not null" json:"link_aplikasi"`
	SertifikasiAplikasi     string    `gorm:"column:sertifikasi_aplikasi" json:"sertifikasi_aplikasi"`
	TanggalAktif            time.Time `gorm:"column:tanggal_aktif;not null" json:"tanggal_aktif"`
	TanggalKadaluarsa       time.Time `gorm:"column:tanggal_kadaluarsa;not null" json:"tanggal_kadaluarsa"`
	AssetID                 int64     `gorm:"column:asset_id;not null" json:"asset_id"`
	Asset                   Asset     `gorm:"foreignKey:AssetID;references:ID" json:"asset"`
	CreatedAt               time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (*DetaiAsetAplikasi) TableName() string {
	return TableNameDetaiAsetAplikasi
}
