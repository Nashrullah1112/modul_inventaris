package Repository

import (
	"itam/Model/Database"
	"time"

	"gorm.io/gorm"
)

type (
	AssetAplikasiRepositoryHandler interface {
		Save(data *Database.DetaiAsetAplikasi) (id int64, err error)
		Update(data *Database.DetaiAsetAplikasi) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetaiAsetAplikasi, err error)
		FindAll() (data []Database.DetaiAsetAplikasi, err error)
		TotalAplikasi(status string) (total int64, err error)
	}

	AssetAplikasiRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetAplikasiRepositoryProvider(db *gorm.DB) *AssetAplikasiRepositoryImpl {
	return &AssetAplikasiRepositoryImpl{
		DB: db,
	}
}

func (h *AssetAplikasiRepositoryImpl) Save(data *Database.DetaiAsetAplikasi) (id int64, err error) {
	err = h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetAplikasiRepositoryImpl) Update(data *Database.DetaiAsetAplikasi) (id int64, err error) {
	err = h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetAplikasiRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetAplikasiRepositoryImpl) FindById(id int64) (data Database.DetaiAsetAplikasi, err error) {
	err = h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetAplikasiRepositoryImpl) FindAll() (data []Database.DetaiAsetAplikasi, err error) {
	err = h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
func (h *AssetAplikasiRepositoryImpl) TotalAplikasi(status string) (total int64, err error) {

	err = h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Joins("JOIN assets ON assets.id = detai_aset_aplikasi.asset_id").
		Where("assets.status != ?", status).
		Count(&total).
		Error
	return total, err
}
