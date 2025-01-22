package Services

import (
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"net/http"
)

type (
	AssetAplikasiServiceHandler interface {
		Create(request Response.DetaiAsetAplikasiCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(aplikasiID int64, request Response.DetaiAsetAplikasiUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(detaiAsetAplikasiId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(detaiAsetAplikasiId int64) (detaiAsetAplikasi Response.DetaiAsetAplikasiResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (detaiAsetAplikasis []Response.DetaiAsetAplikasiResponse, serviceErr *Web.ServiceErrorDto)
		TotalAplikasi() (total int64, err *Web.ServiceErrorDto)
	}

	AssetAplikasiServiceImpl struct {
		aplikasiRepo Repository.AssetAplikasiRepositoryHandler
		assetRepo    Repository.AssetRepositoryHandler
	}
)

func AssetAplikasiServiceProvider(aplikasiRepo Repository.AssetAplikasiRepositoryHandler, assetRepo Repository.AssetRepositoryHandler) *AssetAplikasiServiceImpl {
	return &AssetAplikasiServiceImpl{
		aplikasiRepo: aplikasiRepo,
		assetRepo:    assetRepo,
	}
}

func (h *AssetAplikasiServiceImpl) Create(request Response.DetaiAsetAplikasiCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	assetId, err := h.assetRepo.Save(&Database.Asset{
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Merk:         request.Merk,
		NomorNota:    request.NomorNota,
		VendorID:     request.VendorID,
		Status:       "Approval",
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Aset Hardware not created", err, http.StatusInternalServerError)
	}
	if id, err := h.aplikasiRepo.Save(&Database.DetaiAsetAplikasi{
		NamaAplikasi:            request.NamaAplikasi,
		TanggalPembuatan:        request.TanggalPembuatan,
		TanggalTerima:           request.TanggalTerima,
		LokasiServerPenyimpanan: request.LokasiServerPenyimpanan,
		TipeAplikasi:            request.TipeAplikasi,
		LinkAplikasi:            request.LinkAplikasi,
		SertifikasiAplikasi:     request.SertifikasiAplikasi,
		TanggalAktif:            request.TanggalAktif,
		TanggalKadaluarsa:       request.TanggalKadaluarsa,
		AssetID:                 assetId,
	}); err != nil {
		return id, Web.NewCustomServiceError("DetaiAsetAplikasi not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *AssetAplikasiServiceImpl) Update(aplikasiID int64, request Response.DetaiAsetAplikasiUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	// Find the DetaiAsetAplikasi by ID
	existingDetaiAsetAplikasi, err := h.aplikasiRepo.FindById(aplikasiID)
	if err != nil {
		return 0, Web.NewCustomServiceError("DetaiAsetAplikasi not found", err, http.StatusNotFound)
	}
	if id, err := h.assetRepo.Update(&Database.Asset{
		ID:           existingDetaiAsetAplikasi.AssetID,
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Merk:         request.Merk,
		NomorNota:    request.NomorNota,
		VendorID:     request.VendorID,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}
	if id, err := h.aplikasiRepo.Update(&Database.DetaiAsetAplikasi{
		ID:                      existingDetaiAsetAplikasi.ID,
		NamaAplikasi:            request.NamaAplikasi,
		TanggalPembuatan:        request.TanggalPembuatan,
		TanggalTerima:           request.TanggalTerima,
		LokasiServerPenyimpanan: request.LokasiServerPenyimpanan,
		TipeAplikasi:            request.TipeAplikasi,
		LinkAplikasi:            request.LinkAplikasi,
		SertifikasiAplikasi:     request.SertifikasiAplikasi,
		TanggalAktif:            request.TanggalAktif,
		TanggalKadaluarsa:       request.TanggalKadaluarsa,
		AssetID:                 existingDetaiAsetAplikasi.AssetID,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *AssetAplikasiServiceImpl) Delete(detaiAsetAplikasiId int64) (serviceErr *Web.ServiceErrorDto) {
	// Check if the DetaiAsetAplikasi exists
	asset, err := h.aplikasiRepo.FindById(detaiAsetAplikasiId)
	if err != nil {
		return Web.NewCustomServiceError("DetaiAsetAplikasi not found", err, http.StatusNotFound)
	}

	if err := h.assetRepo.Delete(asset.AssetID); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *AssetAplikasiServiceImpl) FindById(detaiAsetAplikasiId int64) (detaiAsetAplikasi Response.DetaiAsetAplikasiResponse, serviceErr *Web.ServiceErrorDto) {
	// Retrieve the DetaiAsetAplikasi from the database
	data, err := h.aplikasiRepo.FindById(detaiAsetAplikasiId)
	if err != nil {
		return Response.DetaiAsetAplikasiResponse{}, Web.NewCustomServiceError("DetaiAsetAplikasi not found", err, http.StatusNotFound)
	}
	// Convert Database.DetaiAsetAplikasi to Response.DetaiAsetAplikasiResponse
	detaiAsetAplikasi = Response.DetaiAsetAplikasiResponse{
		Id:                      data.ID,
		NamaAplikasi:            data.NamaAplikasi,
		TanggalPembuatan:        data.TanggalPembuatan,
		TanggalTerima:           data.TanggalTerima,
		LokasiServerPenyimpanan: data.LokasiServerPenyimpanan,
		TipeAplikasi:            data.TipeAplikasi,
		LinkAplikasi:            data.LinkAplikasi,
		SertifikasiAplikasi:     data.SertifikasiAplikasi,
		TanggalAktif:            data.TanggalAktif,
		TanggalKadaluarsa:       data.TanggalKadaluarsa,
		AssetID:                 data.AssetID,
		Asset: Response.AssetResponse{
			Id:           data.Asset.ID,
			SerialNumber: data.Asset.SerialNumber,
			Model:        data.Asset.Model,
			Merk:         data.Asset.Merk,
			NomorNota:    data.Asset.NomorNota,
			VendorID:     data.Asset.VendorID,
			Vendor: Response.VendorResponse{
				ID:               data.Asset.Vendor.ID,
				PIC:              data.Asset.Vendor.PIC,
				NomorKontak:      data.Asset.Vendor.NomorKontak,
				LokasiPerusahaan: data.Asset.Vendor.Lokasi,
				NomorSIUP:        data.Asset.Vendor.NomorSIUP,
				NomorNIB:         data.Asset.Vendor.NomorNIB,
				NomorNPWP:        data.Asset.Vendor.NomorNPWP,
			},
		},
	}

	return detaiAsetAplikasi, nil
}

func (h *AssetAplikasiServiceImpl) FindAll() (detaiAsetAplikasis []Response.DetaiAsetAplikasiResponse, serviceErr *Web.ServiceErrorDto) {
	// Retrieve all DetaiAsetAplikasis from the database
	data, err := h.aplikasiRepo.FindAll()
	if err != nil {
		return []Response.DetaiAsetAplikasiResponse{}, Web.NewInternalServiceError(err)
	}

	// Convert each Database.DetaiAsetAplikasi to Response.DetaiAsetAplikasiResponse
	for _, d := range data {

		if d.Asset.Status == "Approved" {
			detaiAsetAplikasis = append(detaiAsetAplikasis, Response.DetaiAsetAplikasiResponse{
				Id:                      d.ID,
				NamaAplikasi:            d.NamaAplikasi,
				TanggalPembuatan:        d.TanggalPembuatan,
				TanggalTerima:           d.TanggalTerima,
				LokasiServerPenyimpanan: d.LokasiServerPenyimpanan,
				TipeAplikasi:            d.TipeAplikasi,
				LinkAplikasi:            d.LinkAplikasi,
				SertifikasiAplikasi:     d.SertifikasiAplikasi,
				TanggalAktif:            d.TanggalAktif,
				TanggalKadaluarsa:       d.TanggalKadaluarsa,
				AssetID:                 d.AssetID,
				Asset: Response.AssetResponse{
					Id:           d.Asset.ID,
					SerialNumber: d.Asset.SerialNumber,
					Model:        d.Asset.Model,
					Merk:         d.Asset.Merk,
					NomorNota:    d.Asset.NomorNota,
					VendorID:     d.Asset.Vendor.ID,
					Vendor: Response.VendorResponse{
						ID:               d.Asset.Vendor.ID,
						PIC:              d.Asset.Vendor.PIC,
						Email:            d.Asset.Vendor.Email,
						NomorKontak:      d.Asset.Vendor.NomorKontak,
						LokasiPerusahaan: d.Asset.Vendor.Lokasi,
						NomorSIUP:        d.Asset.Vendor.NomorSIUP,
						NomorNIB:         d.Asset.Vendor.NomorNIB,
						NomorNPWP:        d.Asset.Vendor.NomorNPWP,
					},
				},
			})
		}

	}

	return detaiAsetAplikasis, nil
}

func (h *AssetAplikasiServiceImpl) TotalAplikasi() (total int64, serviceErr *Web.ServiceErrorDto) {
	total, err := h.aplikasiRepo.TotalAplikasi("Disposal")
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}
	return total, nil
}
