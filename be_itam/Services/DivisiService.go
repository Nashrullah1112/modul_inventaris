package Services

import (
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"net/http"
)

type (
	DivisiServiceHandler interface {
		Create(request Response.DivisiCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.DivisiUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(divisiId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(divisiId int64) (divisi Response.DivisiResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (divisi []Response.DivisiResponse, serviceErr *Web.ServiceErrorDto)
	}

	DivisiServiceImpl struct {
		repo Repository.DivisiRepositoryHandler
	}
)

func DivisiServiceProvider(repo Repository.DivisiRepositoryHandler) *DivisiServiceImpl {
	return &DivisiServiceImpl{
		repo: repo,
	}
}

// Create Divisi
func (h *DivisiServiceImpl) Create(request Response.DivisiCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	id, err := h.repo.Save(&Database.Divisi{
		Nama: request.Nama,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Divisi not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

// Update Divisi
func (h *DivisiServiceImpl) Update(request Response.DivisiUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingDivisi, err := h.repo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Divisi not found", err, http.StatusNotFound)
	}

	id, err = h.repo.Update(&Database.Divisi{
		ID:   existingDivisi.ID,
		Nama: request.Nama,
	})
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}

	return id, nil
}

// Delete Divisi
func (h *DivisiServiceImpl) Delete(divisiId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(divisiId)
	if err != nil {
		return Web.NewCustomServiceError("Divisi not found", err, http.StatusNotFound)
	}

	if err := h.repo.Delete(divisiId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

// Find Divisi by ID
func (h *DivisiServiceImpl) FindById(divisiId int64) (divisi Response.DivisiResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindById(divisiId)
	if err != nil {
		return Response.DivisiResponse{}, Web.NewCustomServiceError("Divisi not found", err, http.StatusNotFound)
	}

	divisi = Response.DivisiResponse{
		ID:   data.ID,
		Nama: data.Nama,
	}

	return divisi, nil
}

// Find All Divisi
func (h *DivisiServiceImpl) FindAll() (divisi []Response.DivisiResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindAll()
	if err != nil {
		return []Response.DivisiResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		divisi = append(divisi, Response.DivisiResponse{
			ID:   d.ID,
			Nama: d.Nama,
		})
	}

	return divisi, nil
}
