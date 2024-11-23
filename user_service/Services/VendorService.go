package Services

import (
	"net/http"
	"user/Model/Database"
	"user/Model/Web"
	"user/Model/Web/Response"
	"user/Repository"
)

type (
	VendorServiceHandler interface {
		Create(request Response.VendorCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.VendorUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(vendorId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(vendorId int64) (vendor Response.VendorResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (vendors []Response.VendorResponse, serviceErr *Web.ServiceErrorDto)
	}

	VendorServiceImpl struct {
		repo Repository.VendorRepositoryHandler
	}
)

func VendorServiceProvider(repo Repository.VendorRepositoryHandler) *VendorServiceImpl {
	return &VendorServiceImpl{
		repo: repo,
	}
}

func (h *VendorServiceImpl) Create(request Response.VendorCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	id, err := h.repo.Save(&Database.Vendor{
		Nama: request.Nama,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Vendor not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *VendorServiceImpl) Update(request Response.VendorUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingVendor, err := h.repo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Vendor not found", err, http.StatusNotFound)
	}

	id, err = h.repo.Update(&Database.Vendor{
		ID:   existingVendor.ID,
		Nama: request.Nama,
	})
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *VendorServiceImpl) Delete(vendorId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(vendorId)
	if err != nil {
		return Web.NewCustomServiceError("Vendor not found", err, http.StatusNotFound)
	}

	if err := h.repo.Delete(vendorId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *VendorServiceImpl) FindById(vendorId int64) (vendor Response.VendorResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindById(vendorId)
	if err != nil {
		return Response.VendorResponse{}, Web.NewCustomServiceError("Vendor not found", err, http.StatusNotFound)
	}

	vendor = Response.VendorResponse{
		ID:   data.ID,
		Nama: data.Nama,
	}

	return vendor, nil
}

func (h *VendorServiceImpl) FindAll() (vendors []Response.VendorResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindAll()
	if err != nil {
		return []Response.VendorResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		vendors = append(vendors, Response.VendorResponse{
			ID:   d.ID,
			Nama: d.Nama,
		})
	}

	return vendors, nil
}
