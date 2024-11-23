package Repository

import (
	"user/Model/Database"

	"gorm.io/gorm"
)

type (
	RoleRepositoryHandler interface {
		Save(data *Database.Role) (id int64, err error)
		Update(data *Database.Role) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Role, err error)
		FindAll() (data []Database.Role, err error)
	}

	RoleRepositoryImpl struct {
		DB *gorm.DB
	}
)

func RoleRepositoryProvider(db *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{
		DB: db,
	}
}

func (h *RoleRepositoryImpl) Save(data *Database.Role) (id int64, err error) {
	err = h.DB.Model(&Database.Role{}).
		Save(&data).Error

	return data.ID, err
}

func (h *RoleRepositoryImpl) Update(data *Database.Role) (id int64, err error) {
	err = h.DB.Model(&Database.Role{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *RoleRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.Role{}, id).Error
	return err
}

func (h *RoleRepositoryImpl) FindById(id int64) (data Database.Role, err error) {
	err = h.DB.Model(&Database.Role{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *RoleRepositoryImpl) FindAll() (data []Database.Role, err error) {
	err = h.DB.Model(&Database.Role{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
