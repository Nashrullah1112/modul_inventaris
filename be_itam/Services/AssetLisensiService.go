package Services

import (
	"fmt"
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"log"
	"net/http"
)

type (
	AssetLisensiServiceHandler interface {
		Create(request Response.AssetLicenseCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(lisensiID int64, request Response.AssetLicenseUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(detailAsetLisensiId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(detailAsetLisensiId int64) (detailAsetLisensi Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (detailAsetLisensi []Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto)
		TotalLisensi() (total int64, serviceErr *Web.ServiceErrorDto)
		NotifyExpiringLicenses() (notifications []Database.Notification, serviceErr *Web.ServiceErrorDto)
		MarkNotificationAsRead(notificationID int64) (serviceErr *Web.ServiceErrorDto)
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
	log.Println("request", request)
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

func (h *AssetLisensiServiceImpl) Update(lisensiID int64, request Response.AssetLicenseUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingDetailAsetLisensi, err := h.lisensiRepo.FindById(lisensiID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Detail Aset Lisensi not found", err, http.StatusNotFound)
	}

	_, err = h.assetRepo.Update(&Database.Asset{
		ID:           existingDetailAsetLisensi.AssetID,
		SerialNumber: request.SNPerangkatTerpasang,
		Model:        request.Model,
		Merk:         request.Merk,
		NomorNota:    request.NomorNota,
		VendorID:     request.VendorID,
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("Aset Hardware not updated", err, http.StatusInternalServerError)
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
		AssetID:                  existingDetailAsetLisensi.AssetID,
	}); err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *AssetLisensiServiceImpl) Delete(detailAsetLisensiId int64) (serviceErr *Web.ServiceErrorDto) {
	asset, err := h.lisensiRepo.FindById(detailAsetLisensiId)
	if err != nil {
		return Web.NewCustomServiceError("Detail Aset Lisensi not found", err, http.StatusNotFound)
	}

	if err := h.assetRepo.Delete(asset.AssetID); err != nil {
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
		AssetID:                  data.AssetID,
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
				Email:            data.Asset.Vendor.Email,
				NomorSIUP:        data.Asset.Vendor.NomorSIUP,
				NomorNIB:         data.Asset.Vendor.NomorNIB,
				NomorNPWP:        data.Asset.Vendor.NomorNPWP,
				LokasiPerusahaan: data.Asset.Vendor.Lokasi,
			},
		},
	}

	return detailAsetLisensi, nil
}

func (h *AssetLisensiServiceImpl) FindAll() (detailAsetLisensi []Response.DetailAsetLisensiResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.lisensiRepo.FindAll()
	if err != nil {
		return []Response.DetailAsetLisensiResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		if d.Asset.Status == "Approved" {
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
				AssetID:                  d.AssetID,
				Asset: Response.AssetResponse{
					Id:           d.Asset.ID,
					SerialNumber: d.Asset.SerialNumber,
					Model:        d.Asset.Model,
					Merk:         d.Asset.Merk,
					NomorNota:    d.Asset.NomorNota,
					VendorID:     d.Asset.VendorID,
					Vendor: Response.VendorResponse{
						ID:               d.Asset.Vendor.ID,
						PIC:              d.Asset.Vendor.PIC,
						NomorKontak:      d.Asset.Vendor.NomorKontak,
						Email:            d.Asset.Vendor.Email,
						NomorSIUP:        d.Asset.Vendor.NomorSIUP,
						NomorNIB:         d.Asset.Vendor.NomorNIB,
						NomorNPWP:        d.Asset.Vendor.NomorNPWP,
						LokasiPerusahaan: d.Asset.Vendor.Lokasi,
					},
				},
			})
		}
	}

	return detailAsetLisensi, nil
}

func (h *AssetLisensiServiceImpl) TotalLisensi() (total int64, serviceErr *Web.ServiceErrorDto) {
	total, err := h.lisensiRepo.TotalLisensi("Disposal")
	if err != nil {
		return 0, Web.NewInternalServiceError(err)

	}
	return total, nil
}

func (h *AssetLisensiServiceImpl) NotifyExpiringLicenses() (notifications []Database.Notification, serviceErr *Web.ServiceErrorDto) {
	// Get all licenses that will expire in 30 days
	err := h.lisensiRepo.SyncNotification()
	if err != nil {
		return nil, Web.NewInternalServiceError(err)
	}

	notifications, err = h.lisensiRepo.GetNotifications("")
	if err != nil {
		return nil, Web.NewInternalServiceError(err)
	}

	return notifications, nil
}
func (h *AssetLisensiServiceImpl) MarkNotificationAsRead(notificationID int64) (serviceErr *Web.ServiceErrorDto) {
	err := h.lisensiRepo.MarkNotificationAsRead(notificationID)
	if err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}
