package Route

import (
	"itam/Config"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(c *fiber.App) {

	assetController := assetDI(Config.DB)
	c.Get("/api/asset", assetController.FindAll)
	c.Get("/api/asset/:assetId", assetController.FindById)
	c.Post("/api/asset", assetController.Create)
	c.Patch("/api/asset", assetController.Update)
	c.Delete("/api/asset/:assetId", assetController.Delete)

	// Asset Aplikasi Routes
	assetAplikasiController := AssetAplikasiDI(Config.DB)
	c.Get("/api/asset-aplikasi", assetAplikasiController.FindAll)
	c.Get("/api/asset-aplikasi/:assetAplikasiId", assetAplikasiController.FindById)
	c.Post("/api/asset-aplikasi", assetAplikasiController.Create)
	c.Patch("/api/asset-aplikasi", assetAplikasiController.Update)
	c.Delete("/api/asset-aplikasi/:assetAplikasiId", assetAplikasiController.Delete)

	// Asset Hardware Routes
	assetHardwareController := assetHardwareDI(Config.DB)
	c.Get("/api/asset-hardware", assetHardwareController.FindAll)
	c.Get("/api/asset-hardware/:assetHardwareId", assetHardwareController.FindById)
	c.Post("/api/asset-hardware", assetHardwareController.Create)
	c.Patch("/api/asset-hardware", assetHardwareController.Update)
	c.Delete("/api/asset-hardware/:assetHardwareId", assetHardwareController.Delete)
	c.Post("/api/form-hardware", assetHardwareController.FormAssetHardware)

	// Asset Lisensi Routes
	assetLisensiController := AssetLisensiDI(Config.DB)
	c.Get("/api/asset-lisensi", assetLisensiController.FindAll)
	c.Get("/api/asset-lisensi/:assetLisensiId", assetLisensiController.FindById)
	c.Post("/api/asset-lisensi", assetLisensiController.Create)
	c.Patch("/api/asset-lisensi", assetLisensiController.Update)
	c.Delete("/api/asset-lisensi/:assetLisensiId", assetLisensiController.Delete)

	// Asset Perangkat Routes
	assetPerangkatController := assetPerangkatDI(Config.DB)
	c.Get("/api/asset-perangkat", assetPerangkatController.FindAll)
	c.Get("/api/asset-perangkat/:assetPerangkatId", assetPerangkatController.FindById)
	c.Post("/api/asset-perangkat", assetPerangkatController.Create)
	c.Patch("/api/asset-perangkat", assetPerangkatController.Update)
	c.Delete("/api/asset-perangkat/:assetPerangkatId", assetPerangkatController.Delete)

	// Vendor Routes
	vendorController := vendorDI(Config.DB)
	c.Get("/api/vendor", vendorController.FindAll)
	c.Get("/api/vendor/:vendorId", vendorController.FindById)
	c.Post("/api/vendor", vendorController.Create)
	c.Patch("/api/vendor", vendorController.Update)
	c.Delete("/api/vendor/:vendorId", vendorController.Delete)

	// User Routes
	userController := userDI(Config.DB)
	c.Get("/api/user", userController.FindAll)
	c.Get("/api/user/:userId", userController.FindById)
	c.Post("/api/user", userController.Create)
	c.Patch("/api/user", userController.Update)
	c.Delete("/api/user/:userId", userController.Delete)

	// Role Routes
	roleController := roleDI(Config.DB)
	c.Get("/api/role", roleController.FindAll)
	c.Get("/api/role/:roleId", roleController.FindById)
	c.Post("/api/role", roleController.Create)
	c.Patch("/api/role", roleController.Update)
	c.Delete("/api/role/:roleId", roleController.Delete)

	// Divisi Routes
	divisiController := divisiDI(Config.DB)
	c.Get("/api/divisi", divisiController.FindAll)
	c.Get("/api/divisi/:divisiId", divisiController.FindById)
	c.Post("/api/divisi", divisiController.Create)
	c.Patch("/api/divisi", divisiController.Update)
	c.Delete("/api/divisi/:divisiId", divisiController.Delete)

}
