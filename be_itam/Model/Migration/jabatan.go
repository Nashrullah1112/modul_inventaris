package Migration

import "time"

const TableNameJabatan = "jabatan"

// Jabatan mapped from table <jabatan>
type Jabatan struct {
	ID        int64     `gorm:"column:id ;primaryKey;autoIncrement:true" json:"id"`
	Nama      string    `gorm:"column:nama;not null" json:"nama"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Jabatan's table name
func (*Jabatan) TableName() string {
	return TableNameJabatan
}
