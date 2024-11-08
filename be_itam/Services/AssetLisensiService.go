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
	AssetLisensiServiceHandler interface {
		Create(request Response.AssetLicenseCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.DetailAsetLisensiUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(detailAsetLisensiId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(detailAsetLisensiId int64) (detailAsetLisensi Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (detailAsetLisensi []Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto)
	}

	AssetLisensiServiceImpl struct {
		lisensiRepo Repository.AssetLisensiRepositoryHandler
		assetRepo   Repository.AssetRepositoryHandler
	}
)

func AssetLisensiServiceProvider(lisensiRepo Repository.AssetLisensiRepositoryHandler, assetRepo Repository.AssetRepositoryHandler) *AssetLisensiServiceImpl {
	return &AssetLisensiServiceImpl{
		lisensiRepo: lisensiRepo,
		assetRepo:   assetRepo,
	}
}

func (h *AssetLisensiServiceImpl) Create(request Response.AssetLicenseCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	assetId, err := h.assetRepo.Save(&Database.Asset{
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Merk:         request.Merk,
		NomorNota:    request.NomorNota,
		VendorID:     request.VendorID,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Aset Hardware not created", err, http.StatusInternalServerError)
	}
	fmt.Println("assetId", assetId)
	if id, err := h.lisensiRepo.Save(&Database.DetailAsetLisensi{
		WaktuPembelian:           request.WaktuPembelian,
		SNPerangkatTerpasang:     request.SNPerangkatTerpasang,
		WaktuAktivasi:            request.WaktuAktivasi,
		TanggalExpired:           request.TanggalExpired,
		TipeKepemilikanAset:      request.TipeKepemilikanAset,
		KategoriLisensi:          request.KategoriLisensi,
		VersiLisensi:             request.VersiLisensi,
		MaksimalUserAplikasi:     request.MaksimalUserAplikasi,
		MaksimalPerangkatLisensi: request.MaksimalPerangkatLisensi,
		TipeLisensi:              request.TipeLisensi,
		AssetID:                  assetId,
	}); err != nil {
		return id, Web.NewCustomServiceError("Detail Aset Lisensi not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *AssetLisensiServiceImpl) Update(request Response.DetailAsetLisensiUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingDetailAsetLisensi, err := h.lisensiRepo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Detail Aset Lisensi not found", err, http.StatusNotFound)
	}

	if id, err := h.lisensiRepo.Update(&Database.DetailAsetLisensi{
		ID:                       existingDetailAsetLisensi.ID,
		WaktuPembelian:           request.WaktuPembelian,
		SNPerangkatTerpasang:     request.SNPerangkatTerpasang,
		WaktuAktivasi:            request.WaktuAktivasi,
		TanggalExpired:           request.TanggalExpired,
		TipeKepemilikanAset:      request.TipeKepemilikanAset,
		KategoriLisensi:          request.KategoriLisensi,
		VersiLisensi:             request.VersiLisensi,
		MaksimalUserAplikasi:     request.MaksimalUserAplikasi,
		MaksimalPerangkatLisensi: request.MaksimalPerangkatLisensi,
		TipeLisensi:              request.TipeLisensi,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *AssetLisensiServiceImpl) Delete(detailAsetLisensiId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.lisensiRepo.FindById(detailAsetLisensiId)
	if err != nil {
		return Web.NewCustomServiceError("Detail Aset Lisensi not found", err, http.StatusNotFound)
	}

	if err := h.lisensiRepo.Delete(detailAsetLisensiId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *AssetLisensiServiceImpl) FindById(detailAsetLisensiId int64) (detailAsetLisensi Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.lisensiRepo.FindById(detailAsetLisensiId)
	if err != nil {
		return Response.DetailAsetLisensiResponse{}, Web.NewCustomServiceError("Detail Aset Lisensi not found", err, http.StatusNotFound)
	}

	detailAsetLisensi = Response.DetailAsetLisensiResponse{
		ID:                       data.ID,
		WaktuPembelian:           data.WaktuPembelian,
		SNPerangkatTerpasang:     data.SNPerangkatTerpasang,
		WaktuAktivasi:            data.WaktuAktivasi,
		TanggalExpired:           data.TanggalExpired,
		TipeKepemilikanAset:      data.TipeKepemilikanAset,
		KategoriLisensi:          data.KategoriLisensi,
		VersiLisensi:             data.VersiLisensi,
		MaksimalUserAplikasi:     data.MaksimalUserAplikasi,
		MaksimalPerangkatLisensi: data.MaksimalPerangkatLisensi,
		TipeLisensi:              data.TipeLisensi,
	}

	return detailAsetLisensi, nil
}

func (h *AssetLisensiServiceImpl) FindAll() (detailAsetLisensi []Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.lisensiRepo.FindAll()
	if err != nil {
		return []Response.DetailAsetLisensiResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		detailAsetLisensi = append(detailAsetLisensi, Response.DetailAsetLisensiResponse{
			ID:                       d.ID,
			WaktuPembelian:           d.WaktuPembelian,
			SNPerangkatTerpasang:     d.SNPerangkatTerpasang,
			WaktuAktivasi:            d.WaktuAktivasi,
			TanggalExpired:           d.TanggalExpired,
			TipeKepemilikanAset:      d.TipeKepemilikanAset,
			KategoriLisensi:          d.KategoriLisensi,
			VersiLisensi:             d.VersiLisensi,
			MaksimalUserAplikasi:     d.MaksimalUserAplikasi,
			MaksimalPerangkatLisensi: d.MaksimalPerangkatLisensi,
			TipeLisensi:              d.TipeLisensi,
		})
	}

	return detailAsetLisensi, nil
}
