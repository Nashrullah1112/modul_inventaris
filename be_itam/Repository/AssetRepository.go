package Repository

import (
	"itam/Model/Database"
	"log"
	"time"

	"gorm.io/gorm"
)

type (
	AssetRepositoryHandler interface {
		Save(data *Database.Asset) (id int64, err error)
		Update(data *Database.Asset) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Asset, err error)
		FindAll() (data []Database.Asset, err error)
		FindDisposal() (data []Database.Asset, err error)
	}

	AssetRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetRepositoryProvider(db *gorm.DB) *AssetRepositoryImpl {
	return &AssetRepositoryImpl{
		DB: db,
	}
}

func (h *AssetRepositoryImpl) Save(data *Database.Asset) (id int64, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetRepositoryImpl) Update(data *Database.Asset) (id int64, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.Asset{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetRepositoryImpl) FindById(id int64) (data Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetRepositoryImpl) FindAll() (data []Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Find(&data).
		Error

	return data, err
}
func (h *AssetRepositoryImpl) FindDisposal() (data []Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).Joins("Vendor").
		Where("status = ?", "Disposal").
		Find(&data).
		Error
	log.Println(data)
	return data, err
}
