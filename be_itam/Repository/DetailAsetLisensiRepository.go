package Repository

import (
	"itam/Model/Database"

	"gorm.io/gorm"
)

type (
	DetailAsetLisensiRepositoryHandler interface {
		Save(data *Database.DetailAsetLisensi) (id int64, err error)
		Update(data *Database.DetailAsetLisensi) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetLisensi, err error)
		FindAll() (data []Database.DetailAsetLisensi, err error)
	}

	DetailAsetLisensiRepositoryImpl struct {
		DB *gorm.DB
	}
)

func DetailAsetLisensiRepositoryProvider(db *gorm.DB) *DetailAsetLisensiRepositoryImpl {
	return &DetailAsetLisensiRepositoryImpl{
		DB: db,
	}
}

func (h *DetailAsetLisensiRepositoryImpl) Save(data *Database.DetailAsetLisensi) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Save(&data).Error

	return data.ID, err
}

func (h *DetailAsetLisensiRepositoryImpl) Update(data *Database.DetailAsetLisensi) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *DetailAsetLisensiRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.DetailAsetLisensi{}, id).Error
	return err
}

func (h *DetailAsetLisensiRepositoryImpl) FindById(id int64) (data Database.DetailAsetLisensi, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *DetailAsetLisensiRepositoryImpl) FindAll() (data []Database.DetailAsetLisensi, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
