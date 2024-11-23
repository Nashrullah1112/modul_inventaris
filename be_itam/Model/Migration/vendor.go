package Migration

import "time"
const TableNameVendor = "vendor"

// Vendor mapped from table <vendor>
type Vendor struct {
	ID   int64  `gorm:"column:id;primaryKey" json:"id"`
	Nama string `gorm:"column:nama;not null" json:"nama"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Vendor's table name
func (*Vendor) TableName() string {
	return TableNameVendor
}
