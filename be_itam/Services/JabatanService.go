package Services

import (
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"log"
	"net/http"
)

type (
	JabatanServiceHandler interface {
		Create(request Response.JabatanCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.JabatanUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(jabatanId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(jabatanId int64) (jabatan Response.JabatanResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (jabatan []Response.JabatanResponse, serviceErr *Web.ServiceErrorDto)
	}

	JabatanServiceImpl struct {
		repo Repository.JabatanRepositoryHandler
	}
)

func JabatanServiceControllerProvider(repo Repository.JabatanRepositoryHandler) *JabatanServiceImpl {
	return &JabatanServiceImpl{
		repo: repo,
	}
}

func (h *JabatanServiceImpl) Create(request Response.JabatanCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	id, err := h.repo.Save(&Database.Jabatan{
		Nama: request.Nama,
	})

	if err != nil {
		log.Println(err)
		return id, Web.NewCustomServiceError("Jabatan not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *JabatanServiceImpl) Update(request Response.JabatanUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingJabatan, err := h.repo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Jabatan not found", err, http.StatusNotFound)
	}

	id, err = h.repo.Update(&Database.Jabatan{
		ID:   existingJabatan.ID,
		Nama: request.Nama,
	})
	if err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *JabatanServiceImpl) Delete(jabatanId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(jabatanId)
	if err != nil {
		return Web.NewCustomServiceError("Jabatan not found", err, http.StatusNotFound)
	}

	if err := h.repo.Delete(jabatanId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *JabatanServiceImpl) FindById(jabatanId int64) (jabatan Response.JabatanResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindById(jabatanId)
	if err != nil {
		return Response.JabatanResponse{}, Web.NewCustomServiceError("Jabatan not found", err, http.StatusNotFound)
	}

	jabatan = Response.JabatanResponse{
		ID:   data.ID,
		Nama: data.Nama,
	}

	return jabatan, nil
}

func (h *JabatanServiceImpl) FindAll() (jabatan []Response.JabatanResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindAll()
	if err != nil {
		return []Response.JabatanResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		jabatan = append(jabatan, Response.JabatanResponse{
			ID:   d.ID,
			Nama: d.Nama,
		})
	}

	return jabatan, nil
}
