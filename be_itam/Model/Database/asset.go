package Database

import (
	"time"
)

type Asset struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	VendorID     int64     `gorm:"column:vendor_id;not null" json:"vendor_id"`
	Vendor       Vendor    `gorm:"foreignKey:VendorID;references:ID" json:"vendor"`
	SerialNumber string    `gorm:"column:serial_number;not null" json:"serial_number"`
	Merk         string    `gorm:"column:merk;not null" json:"merk"`
	Model        string    `gorm:"column:model;not null" json:"model"`
	NomorNota    string    `gorm:"column:nomor_nota;not null" json:"nomor_nota"`
	Status       string    `gorm:"column:status;not null" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
}

func (*Asset) TableName() string {
	return "assets"
}
