package Repository

import (
	"itam/Model/Database"
	"itam/Model/Web/Response"

	"gorm.io/gorm"
)

type (
	RoleRepositoryHandler interface {
		SaveRole(data *Database.Role) (id int64, err error)
		UpdateRole(data *Database.Role) (id int64, err error)
		DeleteRole(id int64) error
		FindRoleById(id int64) (data Database.Role, err error)
		FindRoleByNama(name string) (data Database.Role, err error)
		FindAllRoles() (data []Database.Role, err error)
		SaveModule(data *Database.Module) (id int64, err error)
		UpdateModule(data *Database.Module) (id int64, err error)
		DeleteModule(id int64) error
		FindModuleById(id int64) (data Database.Module, err error)
		FindModuleByNama(name string) (data Database.Module, err error)
		FindAllModules() (data []Database.Module, err error)
		FindAllModulesByRole(id int64) (data []Response.ModuleResponse, err error)
		SaveRoleModule(data *Database.RoleModule) (id int64, err error)
		CheckRoleModuleExist(roleId, moduleId int64) (bool, error)
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

func (h *RoleRepositoryImpl) SaveRole(data *Database.Role) (id int64, err error) {
	err = h.DB.Model(&Database.Role{}).
		Save(&data).Error

	return data.ID, err
}

func (h *RoleRepositoryImpl) UpdateRole(data *Database.Role) (id int64, err error) {
	err = h.DB.Model(&Database.Role{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *RoleRepositoryImpl) DeleteRole(id int64) error {
	err := h.DB.Delete(&Database.Role{}, id).Error
	return err
}

func (h *RoleRepositoryImpl) FindRoleById(id int64) (data Database.Role, err error) {
	err = h.DB.Model(&Database.Role{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *RoleRepositoryImpl) FindAllRoles() (data []Database.Role, err error) {
	err = h.DB.Model(&Database.Role{}).
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
func (h *RoleRepositoryImpl) FindRoleByNama(name string) (data Database.Role, err error) {
	err = h.DB.Model(&Database.Role{}).
		Where("nama = ?", name).
		Take(&data).
		Error
	return data, err
}
func (h *RoleRepositoryImpl) SaveModule(data *Database.Module) (id int64, err error) {
	err = h.DB.Model(&Database.Module{}).Save(&data).Error
	return data.ID, err
}

func (h *RoleRepositoryImpl) UpdateModule(data *Database.Module) (id int64, err error) {
	err = h.DB.Model(&Database.Module{}).
		Where("id = ?", data.ID).
		Updates(&data).Error
	return data.ID, err
}

func (h *RoleRepositoryImpl) DeleteModule(id int64) error {
	return h.DB.Delete(&Database.Module{}, id).Error
}

func (h *RoleRepositoryImpl) FindModuleById(id int64) (data Database.Module, err error) {
	err = h.DB.Model(&Database.Module{}).
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *RoleRepositoryImpl) FindAllModules() (data []Database.Module, err error) {
	err = h.DB.Model(&Database.Module{}).
		Order("id asc").
		Find(&data).
		Error
	return data, err
}
func (h *RoleRepositoryImpl) FindAllModulesByRole(id int64) (data []Response.ModuleResponse, err error) {
	var roleModules []Database.RoleModule
	err = h.DB.Where("role_id = ?", id).Find(&roleModules).Error
	if err != nil {
		return nil, err
	}

	// Ambil semua module IDs dari roleModules
	moduleIDMap := make(map[int64]bool)
	for _, roleModule := range roleModules {
		moduleIDMap[roleModule.ModuleID] = roleModule.IsAllowed
	}

	var modules []Database.Module
	err = h.DB.Order("id asc").Find(&modules).Error
	if err != nil {
		return nil, err
	}

	// Mapping modul ke dalam response dengan status IsAllowed
	var modulesResponse []Response.ModuleResponse
	for _, module := range modules {
		isAllowed := false
		if allowed, exists := moduleIDMap[module.ID]; exists {
			isAllowed = allowed
		}
		modulesResponse = append(modulesResponse, Response.ModuleResponse{
			ID:        module.ID,
			Name:      module.Nama,
			IsAllowed: isAllowed,
		})
	}

	return modulesResponse, nil
}
func (h *RoleRepositoryImpl) FindModuleByNama(name string) (data Database.Module, err error) {
	err = h.DB.Model(&Database.Module{}).
		Where("nama = ?", name).
		Take(&data).
		Error
	return data, err
}

func (h *RoleRepositoryImpl) SaveRoleModule(data *Database.RoleModule) (id int64, err error) {
	err = h.DB.Model(&Database.RoleModule{}).Save(&data).Error
	return data.ID, err
}

func (h *RoleRepositoryImpl) CheckRoleModuleExist(roleId, moduleId int64) (bool, error) {
	var count int64
	err := h.DB.Model(&Database.RoleModule{}).
		Where("role_id = ? AND module_id = ?", roleId, moduleId).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
