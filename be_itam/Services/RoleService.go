package Services

import (
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"net/http"
)

type (
	RoleServiceHandler interface {
		Create(request Response.RoleCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.RoleUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(roleId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(roleId int64) (role Response.RoleResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (roles []Response.RoleResponse, serviceErr *Web.ServiceErrorDto)
	}

	RoleServiceImpl struct {
		repo Repository.RoleRepositoryHandler
	}
)

func RoleServiceControllerProvider(repo Repository.RoleRepositoryHandler) *RoleServiceImpl {
	return &RoleServiceImpl{
		repo: repo,
	}
}

func (h *RoleServiceImpl) Create(request Response.RoleCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	id, err := h.repo.Save(&Database.Role{
		Nama: request.Nama,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Role not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *RoleServiceImpl) Update(request Response.RoleUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingRole, err := h.repo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Role not found", err, http.StatusNotFound)
	}

	id, err = h.repo.Update(&Database.Role{
		ID:   existingRole.ID,
		Nama: request.Nama,
	})
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *RoleServiceImpl) Delete(roleId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(roleId)
	if err != nil {
		return Web.NewCustomServiceError("Role not found", err, http.StatusNotFound)
	}

	if err := h.repo.Delete(roleId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *RoleServiceImpl) FindById(roleId int64) (role Response.RoleResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindById(roleId)
	if err != nil {
		return Response.RoleResponse{}, Web.NewCustomServiceError("Role not found", err, http.StatusNotFound)
	}

	role = Response.RoleResponse{
		ID:   data.ID,
		Nama: data.Nama,
	}

	return role, nil
}

func (h *RoleServiceImpl) FindAll() (roles []Response.RoleResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindAll()
	if err != nil {
		return []Response.RoleResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		roles = append(roles, Response.RoleResponse{
			ID:   d.ID,
			Nama: d.Nama,
		})
	}

	return roles, nil
}
