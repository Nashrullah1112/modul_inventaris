package Services

import (
	"itam/Constant"
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Repository"
	"net/http"
)

type (
	BarangServiceHandler interface {
		Create(request []Web.BarangCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Web.BarangUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(barangId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(barangId int64) (barang Web.BarangResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (barangs []Web.BarangResponse, serviceErr *Web.ServiceErrorDto)
	}

	BarangServiceImpl struct {
		repo Repository.BarangRepositoryHandler
	}
)

func BarangServiceControllerProvider(repo Repository.BarangRepositoryHandler) *BarangServiceImpl {
	return &BarangServiceImpl{
		repo: repo,
	}
}

func (h *BarangServiceImpl) Create(requests []Web.BarangCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	for _, req := range requests {
		if id, err := h.repo.Save(&Database.Barang{
			Name:     req.Name,
			Category: req.Category,
			Price:    req.Price,
		}); err != nil {
			return id, Web.NewCustomServiceError("Barang not create", err, http.StatusNoContent)
		}

	}

	return id, nil
}
func (h *BarangServiceImpl) Update(request Web.BarangUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(request.Id)
	if err != nil {
		return 0, Web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	}

	if id, err := h.repo.Update(&Database.Barang{
		Name:     request.Name,
		Category: request.Category,
		Price:    request.Price,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}
func (h *BarangServiceImpl) Delete(barangID int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(barangID)
	if err != nil {
		return Web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	}

	if err := h.repo.Delete(barangID); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}
func (h *BarangServiceImpl) FindById(barangID int64) (barang Web.BarangResponse, serviceErr *Web.ServiceErrorDto) {
	bar, err := h.repo.FindById(barangID)
	if err != nil {
		return Web.BarangResponse{}, Web.NewCustomServiceError("Barang not found", err, http.StatusNoContent)
	}

	return Constant.ToBarangResponse(bar), nil
}

func (h *BarangServiceImpl) FindAll() (barangs []Web.BarangResponse, serviceErr *Web.ServiceErrorDto) {
	barang, err := h.repo.FindAll()
	if err != nil {
		return []Web.BarangResponse{}, Web.NewInternalServiceError(err)
	}

	return Constant.ToBarangResponses(barang), nil
}
