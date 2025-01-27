package Database

import (
	"time"
)

const TableNameDetailAsetLisensi = "detail_aset_lisensi"

// DetailAsetLisensi mapped from table <detail_aset_lisensi>
type DetailAsetLisensi struct {
	ID                       int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	WaktuPembelian           string    `gorm:"column:waktu_pembelian;not null" json:"waktu_pembelian"`
	SNPerangkatTerpasang     string    `gorm:"column:sn_perangkat_terpasang;not null;unique" json:"sn_perangkat_terpasang"`
	WaktuAktivasi            string    `gorm:"column:waktu_aktivasi;not null" json:"waktu_aktivasi"`
	TanggalExpired           time.Time `gorm:"column:tanggal_expired;not null" json:"tanggal_expired"`
	TipeKepemilikanAset      string    `gorm:"column:tipe_kepemilikan_aset;not null" json:"tipe_kepemilikan_aset"`
	KategoriLisensi          string    `gorm:"column:kategori_lisensi;not null" json:"kategori_lisensi"`
	VersiLisensi             string    `gorm:"column:versi_lisensi;not null" json:"versi_lisensi"`
	MaksimalUserAplikasi     int32     `gorm:"column:maksimal_user_aplikasi;not null" json:"maksimal_user_aplikasi"`
	MaksimalPerangkatLisensi int32     `gorm:"column:maksimal_perangkat_lisensi;not null" json:"maksimal_perangkat_lisensi"`
	TipeLisensi              string    `gorm:"column:tipe_lisensi;not null" json:"tipe_lisensi"`
	AssetID                  int64     `gorm:"column:asset_id;not null" json:"asset_id"`
	Asset                    Asset     `gorm:"foreignKey:AssetID;references:ID" json:"asset"`
	CreatedAt                time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt                time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName DetailAsetLisensi's table name
func (*DetailAsetLisensi) TableName() string {
	return TableNameDetailAsetLisensi
}
