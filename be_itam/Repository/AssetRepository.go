package Repository

import (
	"fmt"
	"itam/Model/Database"
	"itam/Model/Web/Response"
	"log"
	"time"

	"gorm.io/gorm"
)

type (
	AssetRepositoryHandler interface {
		Save(data *Database.Asset) (id int64, err error)
		Update(data *Database.Asset) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.Asset, err error)
		FindAll() (data []Database.Asset, err error)
		FindDisposal() (data []Database.Asset, err error)
		FindApproval() (data []Database.Asset, err error)
		UpdateStatus(id int64, status string) error
		DetailAsset(id int64) (data Response.DetailAssetResponse, err error)
		RejectedAsset(id int64) error
	}

	AssetRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetRepositoryProvider(db *gorm.DB) *AssetRepositoryImpl {
	return &AssetRepositoryImpl{
		DB: db,
	}
}

func (h *AssetRepositoryImpl) Save(data *Database.Asset) (id int64, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Save(&data).Error

	return data.ID, err
}

func (h *AssetRepositoryImpl) Update(data *Database.Asset) (id int64, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.Asset{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetRepositoryImpl) FindById(id int64) (data Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Preload("Vendor").
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetRepositoryImpl) FindAll() (data []Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Find(&data).
		Error

	return data, err
}
func (h *AssetRepositoryImpl) FindDisposal() (data []Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).Joins("Vendor").
		Where("status = ?", "Disposal").
		Find(&data).
		Error
	log.Println(data)
	return data, err
}
func (h *AssetRepositoryImpl) FindApproval() (data []Database.Asset, err error) {
	err = h.DB.Model(&Database.Asset{}).
		Where("status = ? OR status = ?", "Approval", "").
		Find(&data).
		Error
	return data, err
}
func (h *AssetRepositoryImpl) UpdateStatus(id int64, status string) error {
	err := h.DB.Model(&Database.Asset{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": status,
		}).Error
	return err
}
func (h *AssetRepositoryImpl) DetailAsset(id int64) (data Response.DetailAssetResponse, err error) {
	var (
		asset     Database.Asset
		perangkat Database.DetailAsetPerangkat
		aplikasi  Database.DetaiAsetAplikasi
		lisensi   Database.DetailAsetLisensi
		hardware  Database.DetailAsetHardware
		vendor    Database.Vendor
	)
	err = h.DB.Model(&Database.Asset{}).
		Where("id = ?", id).
		Preload("Vendor").
		Find(&asset).
		Error
	if err != nil {
		return data, err
	}
	data.Id = asset.ID
	data.VendorID = asset.VendorID
	data.SerialNumber = asset.SerialNumber
	data.Merk = asset.Merk
	data.Model = asset.Model
	data.NomorNota = asset.NomorNota
	data.Status = asset.Status

	data.Vendor = &Response.VendorResponse{
		ID:               vendor.ID,
		PIC:              vendor.PIC,
		Email:            vendor.Email,
		NomorKontak:      vendor.NomorKontak,
		LokasiPerusahaan: vendor.Lokasi,
		NomorSIUP:        vendor.NomorSIUP,
		NomorNIB:         vendor.NomorNIB,
		NomorNPWP:        vendor.NomorNPWP,
	}

	err = h.DB.Model(&Database.DetailAsetPerangkat{}).
		Where("asset_id = ?", id).
		Find(&perangkat).
		Error
	fmt.Println(perangkat)
	if perangkat.AssetID == id {
		data.Perangkat = &Response.AsetPerangkatResponse{
			ID:                   perangkat.ID,
			LokasiPenerima:       perangkat.LokasiPenerima,
			WaktuPenerimaan:      perangkat.WaktuPenerimaan,
			TandaTerima:          perangkat.TandaTerima,
			TipeAset:             perangkat.TipeAset,
			WaktuAktivasiAset:    perangkat.WaktuAktivasiAset,
			HasilPemeriksaanAset: perangkat.HasilPemeriksaanAset,
			SerialNumber:         perangkat.SerialNumber,
			Model:                perangkat.Model,
			MasaGaransiMulai:     perangkat.MasaGaransiMulai,
			NomorKartuGaransi:    perangkat.NomorKartuGaransi,
			Prosesor:             perangkat.Prosesor,
			KapasitasRAM:         perangkat.KapasitasRAM,
			KapasitasRom:         perangkat.KapasitasRom,
			TipeRAM:              perangkat.TipeRAM,
			TipePenyimpanan:      perangkat.TipePenyimpanan,
			StatusAset:           perangkat.StatusAset,
			NilaiAset:            perangkat.NilaiAset,
			NilaiDepresiasi:      perangkat.NilaiDepresiasi,
			JangkaMasaPakai:      perangkat.JangkaMasaPakai,
			WaktuAsetKeluar:      perangkat.WaktuAsetKeluar,
			KondisiAsetKeluar:    perangkat.KondisiAsetKeluar,
			NotaPembelian:        perangkat.NotaPembelian,
			DivisiID:             perangkat.DivisiID,
			UserID:               perangkat.UserID,
		}

		return data, err
	}
	err = h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("asset_id = ?", id).
		Find(&aplikasi).
		Error
	if aplikasi.AssetID == id {
		data.Aplikasi = &Response.AsetAplikasiResponse{
			Id:                      aplikasi.ID,
			NamaAplikasi:            aplikasi.NamaAplikasi,
			TanggalPembuatan:        aplikasi.TanggalPembuatan,
			TanggalTerima:           aplikasi.TanggalTerima,
			LokasiServerPenyimpanan: aplikasi.LokasiServerPenyimpanan,
			TipeAplikasi:            aplikasi.TipeAplikasi,
			LinkAplikasi:            aplikasi.LinkAplikasi,
			SertifikasiAplikasi:     aplikasi.SertifikasiAplikasi,
			TanggalAktif:            aplikasi.TanggalAktif,
			TanggalKadaluarsa:       aplikasi.TanggalKadaluarsa,
			AssetID:                 aplikasi.AssetID,
		}
		return data, err
	}
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Where("asset_id = ?", id).
		Find(&lisensi).
		Error
	if lisensi.AssetID == id {
		data.Lisensi = &Response.AsetLisensiResponse{
			ID:                       lisensi.ID,
			WaktuPembelian:           lisensi.WaktuPembelian,
			SNPerangkatTerpasang:     lisensi.SNPerangkatTerpasang,
			WaktuAktivasi:            lisensi.WaktuAktivasi,
			TanggalExpired:           lisensi.TanggalExpired,
			TipeKepemilikanAset:      lisensi.TipeKepemilikanAset,
			KategoriLisensi:          lisensi.KategoriLisensi,
			VersiLisensi:             lisensi.VersiLisensi,
			MaksimalUserAplikasi:     lisensi.MaksimalUserAplikasi,
			MaksimalPerangkatLisensi: lisensi.MaksimalPerangkatLisensi,
			TipeLisensi:              lisensi.TipeLisensi,
			AssetID:                  lisensi.AssetID,
		}
		return data, err
	}
	err = h.DB.Model(&Database.DetailAsetHardware{}).
		Where("asset_id = ?", id).
		Find(&hardware).
		Error
	if hardware.AssetID == id {
		data.Hardware = &Response.AsetHardwareResponse{
			Id:                       hardware.ID,
			WaktuPenerimaan:          hardware.WaktuPenerimaan,
			BuktiPenerimaan:          hardware.BuktiPenerimaan,
			TipeAset:                 hardware.TipeAset,
			TanggalAktivasiPerangkat: hardware.TanggalAktivasiPerangkat,
			HasilPemeriksaan:         hardware.HasilPemeriksaan,
			SerialNumber:             hardware.SerialNumber,
			Model:                    hardware.Model,
			WaktuGaransiMulai:        hardware.WaktuGaransiMulai,
			WaktuGaransiBerakhir:     hardware.WaktuGaransiBerakhir,
			NomorKartuGaransi:        hardware.NomorKartuGaransi,
			SpesifikasiPerangkat:     hardware.SpesifikasiPerangkat,
			StatusAset:               hardware.StatusAset,
			PenanggungjawabAset:      hardware.PenanggungjawabAset,
			LokasiPenyimpananID:      hardware.LokasiPenyimpananID,
			JangkaMasaPakai:          hardware.JangkaMasaPakai,
			WaktuAsetKeluar:          hardware.WaktuAsetKeluar,
			KondisiAset:              hardware.KondisiAset,
			NotaPembelian:            hardware.NotaPembelian,
			DivisiID:                 hardware.DivisiID,
			AssetID:                  hardware.AssetID,
		}
		return data, err
	}
	return data, err
}

func (h *AssetRepositoryImpl) RejectedAsset(id int64) error {

	// Hapus data di DetailAsetPerangkat jika ada
	if err := h.DB.Where("asset_id = ?", id).Delete(&Database.DetailAsetPerangkat{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to delete detail perangkat for asset %d: %v", id, err)
	}

	// Hapus data di DetailAsetLisensi jika ada
	if err := h.DB.Where("asset_id = ?", id).Delete(&Database.DetailAsetLisensi{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to delete detail lisensi for asset %d: %v", id, err)
	}

	// Hapus data di DetailAsetAplikasi jika ada
	if err := h.DB.Where("asset_id = ?", id).Delete(&Database.DetaiAsetAplikasi{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to delete detail aplikasi for asset %d: %v", id, err)
	}

	// Hapus data di DetailAsetHardware jika ada
	if err := h.DB.Where("asset_id = ?", id).Delete(&Database.DetailAsetHardware{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to delete detail hardware for asset %d: %v", id, err)
	}

	// Hapus data di tabel Asset
	if err := h.DB.Delete(&Database.Asset{}, id).Error; err != nil && err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to delete asset %d: %v", id, err)
	}

	return nil

}
