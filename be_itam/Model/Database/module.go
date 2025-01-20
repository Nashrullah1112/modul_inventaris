package Database

import "time"

const TableNameModule = "module"

// Role mapped from table <role>
type Module struct {
	ID          int64        `gorm:"column:id;primaryKey" json:"id"`
	Nama        string       `gorm:"column:nama;not null" json:"nama"`
	RoleModules []RoleModule `gorm:"foreignKey:ModuleID"`
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Role's table name
func (*Module) TableName() string {
	return TableNameModule
}
