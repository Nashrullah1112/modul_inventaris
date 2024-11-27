package Repository

import (
	"itam/Model/Database"
	"time"

	"gorm.io/gorm"
)

type (
	AssetHardwareRepositoryHandler interface {
		Save(data *Database.DetailAsetHardware) (id int64, err error)
		Update(data *Database.DetailAsetHardware) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetHardware, err error)
		FindAll() (data []Database.DetailAsetHardware, err error)
		TotalHardware(status string) (total int64, err error)
	}

	AssetHardwareRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetHardwareRepositoryProvider(db *gorm.DB) *AssetHardwareRepositoryImpl {
	return &AssetHardwareRepositoryImpl{
		DB: db,
	}
}

func (h *AssetHardwareRepositoryImpl) Save(data *Database.DetailAsetHardware) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetHardwareRepositoryImpl) Update(data *Database.DetailAsetHardware) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetHardwareRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetHardwareRepositoryImpl) FindById(id int64) (data Database.DetailAsetHardware, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetHardwareRepositoryImpl) FindAll() (data []Database.DetailAsetHardware, err error) {
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}

func (h *AssetHardwareRepositoryImpl) TotalHardware(status string) (total int64, err error) {

	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Joins("JOIN assets ON assets.id = detail_aset_hardware.asset_id").
		Where("assets.status != ?", status).
		Count(&total).
		Error
	return total, err
}
