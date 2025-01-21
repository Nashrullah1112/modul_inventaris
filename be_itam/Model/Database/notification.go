package Database

import (
	"time"
)

// TableNameNotification untuk mendefinisikan nama tabel pada database
const TableNameNotification = "notifications"

// Notification mapped from table <notifications>
type Notification struct {
	ID             int64             `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Message        string            `gorm:"column:message;not null" json:"message"`
	Status         string            `gorm:"column:status;default:'unread';not null" json:"status"`
	AssetLisensiID int64             `gorm:"column:asset_lisensi_id;not null" json:"asset_lisensi_id"`
	AssetLisensi   DetailAsetLisensi `gorm:"foreignKey:AssetLisensiID;references:ID" json:"asset_lisensi"`
	CreatedAt      time.Time         `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time         `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName Notification's table name
func (*Notification) TableName() string {
	return TableNameNotification
}
