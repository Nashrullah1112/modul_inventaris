package Repository

import (
	"itam/Model/Database"

	"gorm.io/gorm"
)

type (
	UserRepositoryHandler interface {
		Save(data *Database.User) (id int64, err error)
		Update(data *Database.User) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.User, err error)
		FindAll() (data []Database.User, err error)
		FindByEmail(email string) (data Database.User, err error)
	}

	UserRepositoryImpl struct {
		DB *gorm.DB
	}
)

func UserRepositoryProvider(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (h *UserRepositoryImpl) Save(data *Database.User) (id int64, err error) {
	err = h.DB.Model(&Database.User{}).
		Save(&data).Error

	return data.ID, err
}

func (h *UserRepositoryImpl) Update(data *Database.User) (id int64, err error) {
	err = h.DB.Model(&Database.User{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *UserRepositoryImpl) Delete(id int64) error {
	err := h.DB.Delete(&Database.User{}, id).Error
	return err
}

func (h *UserRepositoryImpl) FindById(id int64) (data Database.User, err error) {
	err = h.DB.Model(&Database.User{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *UserRepositoryImpl) FindAll() (data []Database.User, err error) {
	err = h.DB.Model(&Database.User{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}

func (h *UserRepositoryImpl) FindByEmail(email string) (data Database.User, err error) {
	err = h.DB.Model(&Database.User{}).
		Where("email = ?", email).
		Take(&data).
		Error
	return data, err
}
