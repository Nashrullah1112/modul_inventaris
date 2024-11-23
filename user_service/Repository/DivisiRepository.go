package Repository

import (
	"user/Model/Database"

	"gorm.io/gorm"
)

type (
	DivisiRepositoryHandler interface {
		Save(data *Database.Divisi) (id int64, err error)
		Update(data *Database.Divisi) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Divisi, err error)
		FindAll() (data []Database.Divisi, err error)
	}

	DivisiRepositoryImpl struct {
		DB *gorm.DB
	}
)

func DivisiRepositoryProvider(db *gorm.DB) *DivisiRepositoryImpl {
	return &DivisiRepositoryImpl{
		DB: db,
	}
}

func (h *DivisiRepositoryImpl) Save(data *Database.Divisi) (id int64, err error) {
	err = h.DB.Model(&Database.Divisi{}).
		Save(&data).Error

	return data.ID, err
}

func (h *DivisiRepositoryImpl) Update(data *Database.Divisi) (id int64, err error) {
	err = h.DB.Model(&Database.Divisi{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *DivisiRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.Divisi{}, id).Error
	return err
}

func (h *DivisiRepositoryImpl) FindById(id int64) (data Database.Divisi, err error) {
	err = h.DB.Model(&Database.Divisi{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *DivisiRepositoryImpl) FindAll() (data []Database.Divisi, err error) {
	err = h.DB.Model(&Database.Divisi{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
