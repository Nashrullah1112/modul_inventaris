package Database

import "time"

const TableNameDivisi = "divisi"

// Divisi mapped from table <divisi>
type Divisi struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nama      string    `gorm:"column:nama;not null" json:"nama"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Divisi's table name
func (*Divisi) TableName() string {
	return TableNameDivisi
}
