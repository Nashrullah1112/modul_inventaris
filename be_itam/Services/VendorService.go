package Services

import (
	"fmt"
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"net/http"
)

type (
	VendorServiceHandler interface {
		Create(request Response.VendorCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(vendorId int64, request Response.VendorUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
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
		PIC:         request.PIC,
		Email:       request.Email,
		NomorKontak: request.NomorKontak,
		Lokasi:      request.LokasiPerusahaan,
		NomorSIUP:   request.NomorSIUP,
		NomorNIB:    request.NomorNIB,
		NomorNPWP:   request.NomorNPWP,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Vendor not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *VendorServiceImpl) Update(vendorId int64, request Response.VendorUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingVendor, err := h.repo.FindById(vendorId)
	if err != nil {
		return 0, Web.NewCustomServiceError("Vendor not found", err, http.StatusNotFound)
	}
	fmt.Println("existingVendor", existingVendor)
	fmt.Println("request", request)
	id, err = h.repo.Update(&Database.Vendor{
		ID:          existingVendor.ID,
		PIC:         request.PIC,
		Email:       request.Email,
		NomorKontak: request.NomorKontak,
		Lokasi:      request.LokasiPerusahaan,
		NomorSIUP:   request.NomorSIUP,
		NomorNIB:    request.NomorNIB,
		NomorNPWP:   request.NomorNPWP,
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
		ID:               data.ID,
		PIC:              data.PIC,
		Email:            data.Email,
		NomorKontak:      data.NomorKontak,
		LokasiPerusahaan: data.Lokasi,
		NomorSIUP:        data.NomorSIUP,
		NomorNIB:         data.NomorNIB,
		NomorNPWP:        data.NomorNPWP,
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
			ID:               d.ID,
			PIC:              d.PIC,
			Email:            d.Email,
			NomorKontak:      d.NomorKontak,
			LokasiPerusahaan: d.Lokasi,
			NomorSIUP:        d.NomorSIUP,
			NomorNIB:         d.NomorNIB,
			NomorNPWP:        d.NomorNPWP,
		})
	}

	return vendors, nil
}
