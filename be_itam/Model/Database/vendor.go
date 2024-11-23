package Database

import "time"

const TableNameVendor = "vendor"

// Vendor mapped from table <vendor>
type Vendor struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	PIC         string    `gorm:"column:pic;not null" json:"pic"`
	Email       string    `gorm:"column:email;not null" json:"email"`
	NomorKontak string    `gorm:"column:nomor_kontak;not null" json:"nomor_kontak"`
	Lokasi      string    `gorm:"column:lokasi;not null" json:"lokasi"`
	NomorSIUP   string    `gorm:"column:nomor_siup;not null" json:"nomor_siup"`
	NomorNIB    string    `gorm:"column:nomor_nib;not null" json:"nomor_nib"`
	NomorNPWP   string    `gorm:"column:nomor_npwp;not null" json:"nomor_npwp"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Vendor's table name
func (*Vendor) TableName() string {
	return TableNameVendor
}
