package Repository

import (
	"user/Model/Database"

	"gorm.io/gorm"
)

type (
	JabatanRepositoryHandler interface {
		Save(data *Database.Jabatan) (id int64, err error)
		Update(data *Database.Jabatan) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Jabatan, err error)
		FindAll() (data []Database.Jabatan, err error)
	}

	JabatanRepositoryImpl struct {
		DB *gorm.DB
	}
)

func JabatanRepositoryProvider(db *gorm.DB) *JabatanRepositoryImpl {
	return &JabatanRepositoryImpl{
		DB: db,
	}
}

func (h *JabatanRepositoryImpl) Save(data *Database.Jabatan) (id int64, err error) {
	err = h.DB.Model(&Database.Jabatan{}).
		Save(&data).Error

	return data.ID, err
}

func (h *JabatanRepositoryImpl) Update(data *Database.Jabatan) (id int64, err error) {
	err = h.DB.Model(&Database.Jabatan{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *JabatanRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.Jabatan{}, id).Error
	return err
}

func (h *JabatanRepositoryImpl) FindById(id int64) (data Database.Jabatan, err error) {
	err = h.DB.Model(&Database.Jabatan{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *JabatanRepositoryImpl) FindAll() (data []Database.Jabatan, err error) {
	err = h.DB.Model(&Database.Jabatan{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
