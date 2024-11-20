package Repository

import (
	"itam/Model/Database"

	"gorm.io/gorm"
)

type (
	AssetAplikasiRepositoryHandler interface {
		Save(data *Database.DetaiAsetAplikasi) (id int64, err error)
		Update(data *Database.DetaiAsetAplikasi) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetaiAsetAplikasi, err error)
		FindAll() (data []Database.DetaiAsetAplikasi, err error)
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
	err := h.DB.Delete(&Database.DetaiAsetAplikasi{}, id).Error
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
