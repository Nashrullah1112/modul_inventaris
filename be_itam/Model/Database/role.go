package Database

import "time"

const TableNameRole = "role"

// Role mapped from table <role>
type Role struct {
	ID          int64        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Nama        string       `gorm:"column:nama;not null" json:"nama"`
	RoleModules []RoleModule `gorm:"foreignKey:RoleID"`
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
