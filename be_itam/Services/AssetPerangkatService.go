package Services

import (
	"itam/Model/Database"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"net/http"
)

type (
	AssetPerangkatServiceHandler interface {
		Create(request Response.DetailAsetPerangkatCreateRequest, tandaTerima string, notaPembelian string) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(perangkatID int64, request Response.DetailAsetPerangkatUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(detailAsetPerangkatId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(detailAsetPerangkatId int64) (detailAsetPerangkat Response.DetailAsetPerangkatResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (detailAsetPerangkat []Response.DetailAsetPerangkatResponse, serviceErr *Web.ServiceErrorDto)
		TotalPerangkat() (total int64, serviceErr *Web.ServiceErrorDto)
	}

	AssetPerangkatServiceImpl struct {
		perangkatRepo Repository.AssetPerangkatRepositoryHandler
		assetRepo     Repository.AssetRepositoryHandler
	}
)

func AssetPerangkatServiceProvider(perangkatRepo Repository.AssetPerangkatRepositoryHandler, assetRepo Repository.AssetRepositoryHandler) *AssetPerangkatServiceImpl {
	return &AssetPerangkatServiceImpl{
		perangkatRepo: perangkatRepo,
		assetRepo:     assetRepo,
	}
}

func (h *AssetPerangkatServiceImpl) Create(request Response.DetailAsetPerangkatCreateRequest, tandaTerima string, notaPembelian string) (id int64, serviceErr *Web.ServiceErrorDto) {
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

	depresiasi := (request.NilaiAset - request.NilaiSisa) / int64(request.JangkaMasaPakai)
	id, err = h.perangkatRepo.Save(&Database.DetailAsetPerangkat{
		LokasiPenerima:       request.LokasiPenerima,
		WaktuPenerimaan:      request.WaktuPenerimaan,
		TandaTerima:          tandaTerima,
		TipeAset:             request.TipeAset,
		WaktuAktivasiAset:    request.WaktuAktivasiAset,
		HasilPemeriksaanAset: request.HasilPemeriksaan,
		SerialNumber:         request.SerialNumber,
		Model:                request.Model,
		MasaGaransiMulai:     request.MasaGaransiMulai,
		NomorKartuGaransi:    request.NomorKartuGaransi,
		Prosesor:             request.Prosesor,
		KapasitasRAM:         request.KapasitasRAM,
		KapasitasRom:         request.KapasitasRom,
		TipeRAM:              request.TipeRAM,
		TipePenyimpanan:      request.TipePenyimpanan,
		StatusAset:           request.StatusAset,
		NilaiAset:            request.NilaiAset,
		NilaiDepresiasi:      depresiasi,
		JangkaMasaPakai:      request.JangkaMasaPakai,
		WaktuAsetKeluar:      request.WaktuAsetKeluar,
		KondisiAsetKeluar:    request.KondisiAsetKeluar,
		NotaPembelian:        notaPembelian,
		DivisiID:             request.DivisiID,
		UserID:               request.UserID,
		AssetID:              assetId,
	})
	if err != nil {
		return id, Web.NewCustomServiceError("Detail Aset Perangkat not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *AssetPerangkatServiceImpl) Update(perangkatID int64, request Response.DetailAsetPerangkatUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingDetailAsetPerangkat, err := h.perangkatRepo.FindById(perangkatID)
	if err != nil {
		return 0, Web.NewCustomServiceError("Detail Aset Perangkat not found", err, http.StatusNotFound)
	}
	_, err = h.assetRepo.Update(&Database.Asset{
		ID:           existingDetailAsetPerangkat.AssetID,
		SerialNumber: request.SerialNumber,
		Model:        request.Model,
		Merk:         request.Merk,
		NomorNota:    request.NomorNota,
		VendorID:     request.VendorID,
	})
	if err != nil {
		return id, Web.NewCustomServiceError("Aset Hardware not updated", err, http.StatusInternalServerError)
	}
	id, err = h.perangkatRepo.Update(&Database.DetailAsetPerangkat{
		ID:                existingDetailAsetPerangkat.ID,
		LokasiPenerima:    request.LokasiPenerima,
		WaktuPenerimaan:   request.WaktuPenerimaan,
		TipeAset:          request.TipeAset,
		WaktuAktivasiAset: request.WaktuAktivasiAset,
		SerialNumber:      request.SerialNumber,
		Model:             request.Model,
		MasaGaransiMulai:  request.MasaGaransiMulai,
		NomorKartuGaransi: request.NomorKartuGaransi,
		Prosesor:          request.Prosesor,
		KapasitasRAM:      request.KapasitasRAM,
		KapasitasRom:      request.KapasitasRom,
		TipeRAM:           request.TipeRAM,
		TipePenyimpanan:   request.TipePenyimpanan,
		StatusAset:        request.StatusAset,
		NilaiAset:         request.NilaiAset,
		NilaiDepresiasi:   request.NilaiDepresiasi,
		JangkaMasaPakai:   request.JangkaMasaPakai,
		WaktuAsetKeluar:   request.WaktuAsetKeluar,
		KondisiAsetKeluar: request.KondisiAsetKeluar,
		DivisiID:          request.DivisiID,
		UserID:            request.UserID,
		AssetID:           existingDetailAsetPerangkat.AssetID,
	})
	if err != nil {
		return id, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *AssetPerangkatServiceImpl) Delete(detailAsetPerangkatId int64) (serviceErr *Web.ServiceErrorDto) {
	asset, err := h.perangkatRepo.FindById(detailAsetPerangkatId)
	if err != nil {
		return Web.NewCustomServiceError("Detail Aset Perangkat not found", err, http.StatusNotFound)
	}

	if err := h.assetRepo.Delete(asset.AssetID); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *AssetPerangkatServiceImpl) FindById(detailAsetPerangkatId int64) (detailAsetPerangkat Response.DetailAsetPerangkatResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.perangkatRepo.FindById(detailAsetPerangkatId)
	if err != nil {
		return Response.DetailAsetPerangkatResponse{}, Web.NewCustomServiceError("Detail Aset Perangkat not found", err, http.StatusNotFound)
	}

	detailAsetPerangkat = Response.DetailAsetPerangkatResponse{
		ID:                   data.ID,
		LokasiPenerima:       data.LokasiPenerima,
		WaktuPenerimaan:      data.WaktuPenerimaan,
		TandaTerima:          data.TandaTerima,
		TipeAset:             data.TipeAset,
		WaktuAktivasiAset:    data.WaktuAktivasiAset,
		HasilPemeriksaanAset: data.HasilPemeriksaanAset,
		SerialNumber:         data.SerialNumber,
		Model:                data.Model,
		MasaGaransiMulai:     data.MasaGaransiMulai,
		NomorKartuGaransi:    data.NomorKartuGaransi,
		Prosesor:             data.Prosesor,
		KapasitasRAM:         data.KapasitasRAM,
		KapasitasRom:         data.KapasitasRom,
		TipeRAM:              data.TipeRAM,
		TipePenyimpanan:      data.TipePenyimpanan,
		StatusAset:           data.StatusAset,
		NilaiAset:            data.NilaiAset,
		NilaiDepresiasi:      data.NilaiDepresiasi,
		JangkaMasaPakai:      data.JangkaMasaPakai,
		WaktuAsetKeluar:      data.WaktuAsetKeluar,
		KondisiAsetKeluar:    data.KondisiAsetKeluar,
		NotaPembelian:        data.NotaPembelian,
		DivisiID:             data.DivisiID,
		UserID:               data.UserID,
		AssetID:              data.AssetID,
		Asset: Response.AssetResponse{
			Id:           data.Asset.ID,
			SerialNumber: data.Asset.SerialNumber,
			Model:        data.Asset.Model,
			Merk:         data.Asset.Merk,
			NomorNota:    data.Asset.NomorNota,
			VendorID:     data.Asset.Vendor.ID,
			Vendor: Response.VendorResponse{
				ID:               data.Asset.Vendor.ID,
				PIC:              data.Asset.Vendor.PIC,
				Email:            data.Asset.Vendor.Email,
				NomorKontak:      data.Asset.Vendor.NomorKontak,
				LokasiPerusahaan: data.Asset.Vendor.Lokasi,
				NomorSIUP:        data.Asset.Vendor.NomorSIUP,
				NomorNIB:         data.Asset.Vendor.NomorNIB,
				NomorNPWP:        data.Asset.Vendor.NomorNPWP,
			},
		},
	}

	return detailAsetPerangkat, nil
}

func (h *AssetPerangkatServiceImpl) FindAll() (detailAsetPerangkat []Response.DetailAsetPerangkatResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.perangkatRepo.FindAll()
	if err != nil {
		return []Response.DetailAsetPerangkatResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		if d.Asset.Status == "Approved" {
			detailAsetPerangkat = append(detailAsetPerangkat, Response.DetailAsetPerangkatResponse{
				ID:                   d.ID,
				LokasiPenerima:       d.LokasiPenerima,
				WaktuPenerimaan:      d.WaktuPenerimaan,
				TandaTerima:          d.TandaTerima,
				TipeAset:             d.TipeAset,
				WaktuAktivasiAset:    d.WaktuAktivasiAset,
				HasilPemeriksaanAset: d.HasilPemeriksaanAset,
				SerialNumber:         d.SerialNumber,
				Model:                d.Model,
				MasaGaransiMulai:     d.MasaGaransiMulai,
				NomorKartuGaransi:    d.NomorKartuGaransi,
				Prosesor:             d.Prosesor,
				KapasitasRAM:         d.KapasitasRAM,
				KapasitasRom:         d.KapasitasRom,
				TipeRAM:              d.TipeRAM,
				TipePenyimpanan:      d.TipePenyimpanan,
				StatusAset:           d.StatusAset,
				NilaiAset:            d.NilaiAset,
				NilaiDepresiasi:      d.NilaiDepresiasi,
				JangkaMasaPakai:      d.JangkaMasaPakai,
				WaktuAsetKeluar:      d.WaktuAsetKeluar,
				KondisiAsetKeluar:    d.KondisiAsetKeluar,
				NotaPembelian:        d.NotaPembelian,
				DivisiID:             d.DivisiID,
				UserID:               d.UserID,
				AssetID:              d.AssetID,
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

	return detailAsetPerangkat, nil
}

func (h *AssetPerangkatServiceImpl) TotalPerangkat() (total int64, serviceErr *Web.ServiceErrorDto) {
	total, err := h.perangkatRepo.TotalPerangkat("Disposal")
	if err != nil {
		return 0, Web.NewInternalServiceError(err)

	}
	return total, nil
}
