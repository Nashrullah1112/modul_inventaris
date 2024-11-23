package Repository

import (
	"itam/Model/Database"

	"gorm.io/gorm"
)

type (
	DetailAsetHardwareRepositoryHandler interface {
		Save(data *Database.DetailAsetHardware) (id int64, err error)
		Update(data *Database.DetailAsetHardware) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetHardware, err error)
		FindAll() (data []Database.DetailAsetHardware, err error)
	}

	DetailAsetHardwareRepositoryImpl struct {
		DB *gorm.DB
	}
)

func DetailAsetHardwareRepositoryProvider(db *gorm.DB) *DetailAsetHardwareRepositoryImpl {
	return &DetailAsetHardwareRepositoryImpl{
		DB: db,
	}
}

func (h *DetailAsetHardwareRepositoryImpl) Save(data *Database.DetailAsetHardware) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Save(&data).Error

	return data.ID, err
}

func (h *DetailAsetHardwareRepositoryImpl) Update(data *Database.DetailAsetHardware) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *DetailAsetHardwareRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.DetailAsetHardware{}, id).Error
	return err
}

func (h *DetailAsetHardwareRepositoryImpl) FindById(id int64) (data Database.DetailAsetHardware, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *DetailAsetHardwareRepositoryImpl) FindAll() (data []Database.DetailAsetHardware, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
