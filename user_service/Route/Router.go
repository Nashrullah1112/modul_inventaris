package Route

import (
	"user/Config"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(c *fiber.App) {

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
