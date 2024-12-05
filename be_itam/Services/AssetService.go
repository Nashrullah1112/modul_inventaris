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
	AssetServiceHandler interface {
		Create(request Response.AssetCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.AssetUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(assetId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(assetId int64) (asset Response.AssetResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (assets []Response.AssetResponse, serviceErr *Web.ServiceErrorDto)
		FindDisposal() (assets []Response.AssetResponse, serviceErr *Web.ServiceErrorDto)
		FindApproval() (assets []Response.AssetResponse, serviceErr *Web.ServiceErrorDto)
		Approval(request Response.IDAsset) (id int64, serviceErr *Web.ServiceErrorDto)
		FindDetailAsset(assetId int64) (detailAsset Response.DetailAssetResponse, serviceErr *Web.ServiceErrorDto)
	}

	AssetServiceImpl struct {
		repo Repository.AssetRepositoryHandler
	}
)

func AssetServiceControllerProvider(repo Repository.AssetRepositoryHandler) *AssetServiceImpl {
	return &AssetServiceImpl{
		repo: repo,
	}
}

func (h *AssetServiceImpl) Create(request Response.AssetCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	// Create asset object from request
	asset := &Database.Asset{
		VendorID:     request.VendorID,
		SerialNumber: request.SerialNumber,
		Merk:         request.Merk,
		Model:        request.Model,
		NomorNota:    request.NomorNota,
	}

	if id, err := h.repo.Save(asset); err != nil {
		return id, Web.NewCustomServiceError("Asset not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *AssetServiceImpl) Update(request Response.AssetUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	// Find the asset by ID
	existingAsset, err := h.repo.FindById(request.Id)
	if err != nil {
		return 0, Web.NewCustomServiceError("Asset not found", err, http.StatusNotFound)
	}

	if id, err := h.repo.Update(&Database.Asset{
		ID:           existingAsset.ID,
		VendorID:     request.VendorID,
		SerialNumber: request.SerialNumber,
		Merk:         request.Merk,
		Model:        request.Model,
		NomorNota:    request.NomorNota,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *AssetServiceImpl) Delete(assetId int64) (serviceErr *Web.ServiceErrorDto) {
	// Check if the asset exists
	_, err := h.repo.FindById(assetId)
	if err != nil {
		return Web.NewCustomServiceError("Asset not found", err, http.StatusNotFound)
	}

	// Delete the asset
	if err := h.repo.Delete(assetId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *AssetServiceImpl) FindById(assetId int64) (asset Response.AssetResponse, serviceErr *Web.ServiceErrorDto) {
	// Retrieve the asset from the database
	data, err := h.repo.FindById(assetId)
	if err != nil {
		return Response.AssetResponse{}, Web.NewCustomServiceError("Asset not found", err, http.StatusNotFound)
	}
	fmt.Println(data)
	// Convert Database.Asset to Response.AssetResponse
	asset = Response.AssetResponse{
		Id:           data.ID,
		VendorID:     data.VendorID,
		SerialNumber: data.SerialNumber,
		Merk:         data.Merk,
		Model:        data.Model,
		NomorNota:    data.NomorNota,
	}

	return asset, nil
}

func (h *AssetServiceImpl) FindAll() (assets []Response.AssetResponse, serviceErr *Web.ServiceErrorDto) {

	data, err := h.repo.FindAll()
	if err != nil {
		return []Response.AssetResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		assets = append(assets, Response.AssetResponse{
			Id:           d.ID,
			VendorID:     d.VendorID,
			SerialNumber: d.SerialNumber,
			Merk:         d.Merk,
			Model:        d.Model,
			NomorNota:    d.NomorNota,
		})
	}

	return assets, nil
}

func (h *AssetServiceImpl) FindDisposal() (assets []Response.AssetResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindDisposal()
	if err != nil {
		return []Response.AssetResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		assets = append(assets, Response.AssetResponse{
			Id:           d.ID,
			VendorID:     d.VendorID,
			SerialNumber: d.SerialNumber,
			Merk:         d.Merk,
			Model:        d.Model,
			NomorNota:    d.NomorNota,
			Vendor: Response.VendorResponse{
				ID:               d.Vendor.ID,
				PIC:              d.Vendor.PIC,
				Email:            d.Vendor.Email,
				NomorKontak:      d.Vendor.NomorKontak,
				LokasiPerusahaan: d.Vendor.Lokasi,
				NomorSIUP:        d.Vendor.NomorSIUP,
				NomorNIB:         d.Vendor.NomorNIB,
				NomorNPWP:        d.Vendor.NomorNPWP,
			},
		})
	}

	return assets, nil

}
func (h *AssetServiceImpl) FindApproval() (assets []Response.AssetResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindApproval()
	if err != nil {
		return []Response.AssetResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		assets = append(assets, Response.AssetResponse{
			Id:           d.ID,
			VendorID:     d.VendorID,
			SerialNumber: d.SerialNumber,
			Merk:         d.Merk,
			Model:        d.Model,
			NomorNota:    d.NomorNota,
			Vendor: Response.VendorResponse{
				ID:               d.Vendor.ID,
				PIC:              d.Vendor.PIC,
				Email:            d.Vendor.Email,
				NomorKontak:      d.Vendor.NomorKontak,
				LokasiPerusahaan: d.Vendor.Lokasi,
				NomorSIUP:        d.Vendor.NomorSIUP,
				NomorNIB:         d.Vendor.NomorNIB,
				NomorNPWP:        d.Vendor.NomorNPWP,
			},
		})
	}

	return assets, nil

}

func (h *AssetServiceImpl) Approval(request Response.IDAsset) (id int64, serviceErr *Web.ServiceErrorDto) {
	// Find the asset by ID
	existingAsset, err := h.repo.FindById(request.Id)
	if err != nil {
		return 0, Web.NewCustomServiceError("Asset not found", err, http.StatusNotFound)
	}

	if err := h.repo.UpdateStatus(existingAsset.ID, "Approved"); err != nil {
		return existingAsset.ID, Web.NewInternalServiceError(err)
	}
	return id, nil
}
func (h *AssetServiceImpl) FindDetailAsset(assetId int64) (detailAsset Response.DetailAssetResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindById(assetId)
	if err != nil {
		return Response.DetailAssetResponse{}, Web.NewCustomServiceError("Detail Asset not found", err, http.StatusNotFound)
	}
	detailAsset, err = h.repo.DetailAsset(data.ID)
	if err != nil {
		return Response.DetailAssetResponse{}, Web.NewInternalServiceError(err)
	}

	return detailAsset, nil
}
