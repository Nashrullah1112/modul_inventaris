package Repository

import (
	"itam/Model/Database"
	"time"

	"gorm.io/gorm"
)

type (
	AssetLisensiRepositoryHandler interface {
		Save(data *Database.DetailAsetLisensi) (id int64, err error)
		Update(data *Database.DetailAsetLisensi) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetLisensi, err error)
		FindAll() (data []Database.DetailAsetLisensi, err error)
		TotalLisensi(status string) (total int64, err error)
	}

	AssetLisensiRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetLisensiRepositoryProvider(db *gorm.DB) *AssetLisensiRepositoryImpl {
	return &AssetLisensiRepositoryImpl{
		DB: db,
	}
}

func (h *AssetLisensiRepositoryImpl) Save(data *Database.DetailAsetLisensi) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetLisensiRepositoryImpl) Update(data *Database.DetailAsetLisensi) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetLisensiRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetLisensiRepositoryImpl) FindById(id int64) (data Database.DetailAsetLisensi, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetLisensiRepositoryImpl) FindAll() (data []Database.DetailAsetLisensi, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
func (h *AssetLisensiRepositoryImpl) TotalLisensi(status string) (total int64, err error) {

	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Joins("JOIN assets ON assets.id = detail_aset_lisensi.asset_id").
		Where("assets.status != ?", status).
		Count(&total).
		Error
	return total, err
}
