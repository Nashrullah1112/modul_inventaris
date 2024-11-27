package Repository

import (
	"itam/Model/Database"
	"time"

	"gorm.io/gorm"
)

type (
	AssetPerangkatRepositoryHandler interface {
		Save(data *Database.DetailAsetPerangkat) (id int64, err error)
		Update(data *Database.DetailAsetPerangkat) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetPerangkat, err error)
		FindAll() (data []Database.DetailAsetPerangkat, err error)
		TotalPerangkat(status string) (total int64, err error)
	}

	AssetPerangkatRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetPerangkatRepositoryProvider(db *gorm.DB) *AssetPerangkatRepositoryImpl {
	return &AssetPerangkatRepositoryImpl{
		DB: db,
	}
}

func (h *AssetPerangkatRepositoryImpl) Save(data *Database.DetailAsetPerangkat) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetPerangkatRepositoryImpl) Update(data *Database.DetailAsetPerangkat) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetPerangkatRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetPerangkatRepositoryImpl) FindById(id int64) (data Database.DetailAsetPerangkat, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetPerangkatRepositoryImpl) FindAll() (data []Database.DetailAsetPerangkat, err error) {
	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}

func (h *AssetPerangkatRepositoryImpl) TotalPerangkat(status string) (total int64, err error) {

	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Joins("JOIN assets ON assets.id = detail_aset_perangkat.asset_id").
		Where("assets.status != ?", status).
		Count(&total).
		Error
	return total, err
}
