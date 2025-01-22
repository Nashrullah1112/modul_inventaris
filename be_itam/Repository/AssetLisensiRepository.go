package Repository

import (
	"fmt"
	"itam/Model/Database"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	AssetLisensiRepositoryHandler interface {
		Save(data *Database.DetailAsetLisensi) (id int64, err error)
		Update(data *Database.DetailAsetLisensi) (id int64, err error)
		Delete(id int64) error
		FindById(id int64) (data Database.DetailAsetLisensi, err error)
		FindAll() (data []Database.DetailAsetLisensi, err error)
		TotalLisensi(status string) (total int64, err error)
		SyncNotification() error
		GetNotifications(status string) ([]Database.Notification, error)
		MarkNotificationAsRead(notificationID int64) error
	}

	AssetLisensiRepositoryImpl struct {
		DB *gorm.DB
	}
)

func AssetLisensiRepositoryProvider(db *gorm.DB) *AssetLisensiRepositoryImpl {
	return &AssetLisensiRepositoryImpl{
		DB: db,
	}
}

func (h *AssetLisensiRepositoryImpl) Save(data *Database.DetailAsetLisensi) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Save(&data).Error

	if err != nil {
		// Periksa apakah error disebabkan oleh constraint unik
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key value") {
			return 0, fmt.Errorf("SN perangkat terpasang '%s' sudah digunakan", data.SNPerangkatTerpasang)
		}
	}
	return data.ID, err
}

func (h *AssetLisensiRepositoryImpl) Update(data *Database.DetailAsetLisensi) (id int64, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Where("id = ?", data.ID).
		Updates(&data).Error

	return data.ID, err
}

func (h *AssetLisensiRepositoryImpl) Delete(id int64) error {
	err := h.DB.Model(&Database.DetaiAsetAplikasi{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_at": time.Now(),
			"status":     "Disposal",
		}).Error
	return err
}

func (h *AssetLisensiRepositoryImpl) FindById(id int64) (data Database.DetailAsetLisensi, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Preload("Asset").
		Preload("Asset.Vendor").
		Where("id = ?", id).
		Take(&data).
		Error
	return data, err
}

func (h *AssetLisensiRepositoryImpl) FindAll() (data []Database.DetailAsetLisensi, err error) {
	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Preload("Asset").
		Preload("Asset.Vendor").
		Order("id asc").
		Find(&data).
		Error

	return data, err
}
func (h *AssetLisensiRepositoryImpl) TotalLisensi(status string) (total int64, err error) {

	err = h.DB.Model(&Database.DetailAsetLisensi{}).
		Joins("JOIN assets ON assets.id = detail_aset_lisensi.asset_id").
		Where("assets.status != ?", status).
		Count(&total).
		Error
	return total, err
}
func (h *AssetLisensiRepositoryImpl) SyncNotification() error {
	// Ambil daftar DetailAsetLisensi yang perlu disinkronkan atau diberi notifikasi
	var assetLisensiList []Database.DetailAsetLisensi
	if err := h.DB.Debug().Preload("Asset").Find(&assetLisensiList).Error; err != nil {
		return fmt.Errorf("failed to get asset lisensi data: %v", err)
	}

	// Iterasi untuk setiap DetailAsetLisensi
	for _, assetLisensi := range assetLisensiList {
		// Cek apakah sudah ada notifikasi dengan status 'unread' atau 'read' untuk asset ini
		var existingNotification Database.Notification
		if err := h.DB.Where("asset_lisensi_id = ?", assetLisensi.ID).First(&existingNotification).Error; err == nil {
			// Jika notifikasi sudah ada, periksa apakah expired telah diperbarui
			if assetLisensi.UpdatedAt.After(existingNotification.UpdatedAt) {
				log.Println("tanggal expired %s > %s", assetLisensi.UpdatedAt, existingNotification.UpdatedAt)
				log.Println(assetLisensi.UpdatedAt.After(existingNotification.UpdatedAt))
				// Perbarui status notifikasi yang lama menjadi expired atau read
				existingNotification.Status = "expired"
				if err := h.DB.Save(&existingNotification).Error; err != nil {
					return fmt.Errorf("failed to update existing notification: %v", err)
				}

				// Buat notifikasi baru dengan status unread
				notification := &Database.Notification{
					Message:        fmt.Sprintf("New notification for asset lisensi: %s with number SN %s", assetLisensi.Asset.Merk, assetLisensi.SNPerangkatTerpasang),
					Status:         "unread", // Status awal adalah unread
					AssetLisensiID: assetLisensi.ID,
				}

				if err := h.DB.Model(&Database.Notification{}).Save(notification).Error; err != nil {
					return fmt.Errorf("failed to create notification: %v", err)
				}
			}
		} else if err == gorm.ErrRecordNotFound {
			// Jika tidak ada notifikasi, buat yang baru
			notification := &Database.Notification{
				Message:        fmt.Sprintf("New notification for asset lisensi: %s with number SN %s", assetLisensi.Asset.Merk, assetLisensi.SNPerangkatTerpasang),
				Status:         "unread", // Status awal adalah unread
				AssetLisensiID: assetLisensi.ID,
			}

			if err := h.DB.Model(&Database.Notification{}).Save(notification).Error; err != nil {
				return fmt.Errorf("failed to create notification: %v", err)
			}
		}
	}

	return nil
}
func (h *AssetLisensiRepositoryImpl) GetNotifications(status string) ([]Database.Notification, error) {
	var notifications []Database.Notification

	// Query untuk mengambil semua notifikasi, jika status diberikan maka memfilter berdasarkan status
	query := h.DB

	// Memfilter berdasarkan status jika status tidak kosong
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Mengurutkan berdasarkan created_at DESC
	if err := query.Order("created_at desc").Find(&notifications).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve notifications: %v", err)
	}

	return notifications, nil
}
func (h *AssetLisensiRepositoryImpl) MarkNotificationAsRead(notificationID int64) error {
	// Mencari notifikasi berdasarkan ID
	var notification Database.Notification
	if err := h.DB.First(&notification, notificationID).Error; err != nil {
		return fmt.Errorf("failed to find notification: %v", err)
	}

	// Memperbarui status notifikasi menjadi 'read'
	notification.Status = "read"
	if err := h.DB.Save(&notification).Error; err != nil {
		return fmt.Errorf("failed to update notification status: %v", err)
	}

	return nil
}
