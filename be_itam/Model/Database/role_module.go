package Database

import "time"

type RoleModule struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	RoleID    int64     `gorm:"column:role_id;not null" json:"role_id"`
	ModuleID  int64     `gorm:"column:module_id;not null" json:"module_id"`
	IsAllowed bool      `gorm:"column:is_allowed;not null" json:"is_allowed"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Role   Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Module Module `gorm:"foreignKey:ModuleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RoleModule) TableName() string {
	return "role_modules"
}
