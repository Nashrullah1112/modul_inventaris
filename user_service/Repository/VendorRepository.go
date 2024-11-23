package Repository

import (
	"user/Model/Database"

	"gorm.io/gorm"
)

type (
	VendorRepositoryHandler interface {
		Save(data *Database.Vendor) (id int64, err error)
		Update(data *Database.Vendor) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Vendor, err error)
		FindAll() (data []Database.Vendor, err error)
	}

	VendorRepositoryImpl struct {
		DB *gorm.DB
	}
)

func VendorRepositoryProvider(db *gorm.DB) *VendorRepositoryImpl {
	return &VendorRepositoryImpl{
		DB: db,
	}
}

func (h *VendorRepositoryImpl) Save(data *Database.Vendor) (id int64, err error) {
	err = h.DB.Model(&Database.Vendor{}).
		Save(&data).Error

	return data.ID, err
}

func (h *VendorRepositoryImpl) Update(data *Database.Vendor) (id int64, err error) {
	err = h.DB.Model(&Database.Vendor{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *VendorRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.Vendor{}, id).Error
	return err
}

func (h *VendorRepositoryImpl) FindById(id int64) (data Database.Vendor, err error) {
	err = h.DB.Model(&Database.Vendor{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *VendorRepositoryImpl) FindAll() (data []Database.Vendor, err error) {
	err = h.DB.Model(&Database.Vendor{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
