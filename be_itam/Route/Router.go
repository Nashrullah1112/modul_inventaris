package Route

import (
	"itam/Config"

	"github.com/gofiber/fiber/v2"
)

// NewRouter sets up the routes for the given Fiber application.
// It initializes various controllers and maps HTTP endpoints to their respective handler functions.
// The following routes are configured:
//
// Asset Routes:
// - GET    /api/asset                : FindAll
// - GET    /api/asset/:assetId       : FindById
// - POST   /api/asset                : Create
// - PATCH  /api/asset                : Update
// - DELETE /api/asset/:assetId       : Delete
//
// Asset Aplikasi Routes:
// - GET    /api/asset-aplikasi                : FindAll
// - GET    /api/asset-aplikasi/:assetAplikasiId : FindById
// - POST   /api/asset-aplikasi                : Create
// - PATCH  /api/asset-aplikasi                : Update
// - DELETE /api/asset-aplikasi/:assetAplikasiId : Delete
//
// Asset Hardware Routes:
// - GET    /api/asset-hardware                : FindAll
// - GET    /api/asset-hardware/:assetHardwareId : FindById
// - POST   /api/asset-hardware                : Create
// - PATCH  /api/asset-hardware                : Update
// - DELETE /api/asset-hardware/:assetHardwareId : Delete
// - POST   /api/form-hardware                 : FormAssetHardware
//
// Asset Lisensi Routes:
// - GET    /api/asset-lisensi                : FindAll
// - GET    /api/asset-lisensi/:assetLisensiId : FindById
// - POST   /api/asset-lisensi                : Create
// - PATCH  /api/asset-lisensi                : Update
// - DELETE /api/asset-lisensi/:assetLisensiId : Delete
//
// Asset Perangkat Routes:
// - GET    /api/asset-perangkat                : FindAll
// - GET    /api/asset-perangkat/:assetPerangkatId : FindById
// - POST   /api/asset-perangkat                : Create
// - PATCH  /api/asset-perangkat                : Update
// - DELETE /api/asset-perangkat/:assetPerangkatId : Delete
//
// Vendor Routes:
// - GET    /api/vendor                : FindAll
// - GET    /api/vendor/:vendorId       : FindById
// - POST   /api/vendor                : Create
// - PATCH  /api/vendor                : Update
// - DELETE /api/vendor/:vendorId       : Delete
//
// User Routes:
// - GET    /api/user                : FindAll
// - GET    /api/user/:userId       : FindById
// - POST   /api/user                : Create
// - PATCH  /api/user                : Update
// - DELETE /api/user/:userId       : Delete
//
// Role Routes:
// - GET    /api/role                : FindAll
// - GET    /api/role/:roleId       : FindById
// - POST   /api/role                : Create
// - PATCH  /api/role                : Update
// - DELETE /api/role/:roleId       : Delete
//
// Divisi Routes:
// - GET    /api/divisi                : FindAll
// - GET    /api/divisi/:divisiId       : FindById
// - POST   /api/divisi                : Create
// - PATCH  /api/divisi                : Update
// - DELETE /api/divisi/:divisiId       : Delete
func NewRouter(c *fiber.App) {

	assetController := AssetDI(Config.DB)
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
	assetHardwareController := AssetHardwareDI(Config.DB)
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
	assetPerangkatController := AssetPerangkatDI(Config.DB)
	c.Get("/api/asset-perangkat", assetPerangkatController.FindAll)
	c.Get("/api/asset-perangkat/:assetPerangkatId", assetPerangkatController.FindById)
	c.Post("/api/asset-perangkat", assetPerangkatController.Create)
	c.Patch("/api/asset-perangkat", assetPerangkatController.Update)
	c.Delete("/api/asset-perangkat/:assetPerangkatId", assetPerangkatController.Delete)

	// Vendor Routes
	vendorController := VendorDI(Config.DB)
	c.Get("/api/vendor", vendorController.FindAll)
	c.Get("/api/vendor/:vendorId", vendorController.FindById)
	c.Post("/api/vendor", vendorController.Create)
	c.Patch("/api/vendor", vendorController.Update)
	c.Delete("/api/vendor/:vendorId", vendorController.Delete)

	// User Routes
	userController := UserDI(Config.DB)
	c.Get("/api/user", userController.FindAll)
	c.Get("/api/user/:userId", userController.FindById)
	c.Post("/api/user", userController.Create)
	c.Patch("/api/user", userController.Update)
	c.Delete("/api/user/:userId", userController.Delete)
	c.Post("/api/login", userController.Login)

	// Role Routes
	roleController := RoleDI(Config.DB)
	c.Get("/api/role", roleController.FindAll)
	c.Get("/api/role/:roleId", roleController.FindById)
	c.Post("/api/role", roleController.Create)
	c.Patch("/api/role", roleController.Update)
	c.Delete("/api/role/:roleId", roleController.Delete)

	// Divisi Routes
	divisiController := DivisiDI(Config.DB)
	c.Get("/api/divisi", divisiController.FindAll)
	c.Get("/api/divisi/:divisiId", divisiController.FindById)
	c.Post("/api/divisi", divisiController.Create)
	c.Patch("/api/divisi", divisiController.Update)
	c.Delete("/api/divisi/:divisiId", divisiController.Delete)

	// Jabatan Routes
	jabatanController := JabatanDI(Config.DB)
	c.Get("/api/jabatan", jabatanController.FindAll)
	c.Get("/api/jabatan/:jabatanId", jabatanController.FindById)
	c.Post("/api/jabatan", jabatanController.Create)
	c.Patch("/api/jabatan", jabatanController.Update)
	c.Delete("/api/jabatan/:jabatanId", jabatanController.Delete)
}
