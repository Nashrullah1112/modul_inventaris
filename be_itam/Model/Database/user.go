package Database

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID               int64     `gorm:"column:id;primaryKey" json:"id"`
	NIP              int64     `gorm:"column:nip;not null" json:"nip"`
	Nama             string    `gorm:"column:nama;not null" json:"nama"`
	Email            string    `gorm:"column:email;not null" json:"email"`
	Password         string    `gorm:"column:password;not null" json:"password"`
	JabatanID        int64     `gorm:"column:jabatan_id;not null" json:"jabatan_id"`
	DivisiID         int64     `gorm:"column:divisi_id;not null" json:"divisi_id"`
	TanggalBergabung time.Time `gorm:"column:tanggal_bergabung;not null" json:"tanggal_bergabung"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Jabatan          Jabatan   `gorm:"foreignKey:JabatanID;references:ID" json:"jabatan"`
	Divisi           Divisi    `gorm:"foreignKey:DivisiID;references:ID" json:"divisi"`
}

// BeforeSave is a GORM hook that hashes the password before saving the user
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
