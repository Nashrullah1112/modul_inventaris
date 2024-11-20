//go:build wireinject
// +build wireinject

package Route

import (
	"user/Controller"
	"user/Repository"
	"user/Services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// Vendor DI
func vendorDI(db *gorm.DB) *Controller.VendorControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.VendorRepositoryProvider,
		Services.VendorServiceProvider,
		Controller.VendorControllerProvider,

		wire.Bind(new(Repository.VendorRepositoryHandler), new(*Repository.VendorRepositoryImpl)),
		wire.Bind(new(Services.VendorServiceHandler), new(*Services.VendorServiceImpl)),
		wire.Bind(new(Controller.VendorControllerHandler), new(*Controller.VendorControllerImpl)),
	)))
	return &Controller.VendorControllerImpl{}
}

// User DI
func userDI(db *gorm.DB) *Controller.UserControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.UserRepositoryProvider,
		Services.UserServiceProvider,
		Controller.UserControllerProvider,

		wire.Bind(new(Repository.UserRepositoryHandler), new(*Repository.UserRepositoryImpl)),
		wire.Bind(new(Services.UserServiceHandler), new(*Services.UserServiceImpl)),
		wire.Bind(new(Controller.UserControllerHandler), new(*Controller.UserControllerImpl)),
	)))
	return &Controller.UserControllerImpl{}
}

// Role DI
func roleDI(db *gorm.DB) *Controller.RoleControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.RoleRepositoryProvider,
		Services.RoleServiceControllerProvider,
		Controller.RoleControllerProvider,

		wire.Bind(new(Repository.RoleRepositoryHandler), new(*Repository.RoleRepositoryImpl)),
		wire.Bind(new(Services.RoleServiceHandler), new(*Services.RoleServiceImpl)),
		wire.Bind(new(Controller.RoleControllerHandler), new(*Controller.RoleControllerImpl)),
	)))
	return &Controller.RoleControllerImpl{}
}

// Divisi DI
func divisiDI(db *gorm.DB) *Controller.DivisiControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.DivisiRepositoryProvider,
		Services.DivisiServiceProvider,
		Controller.DivisiControllerProvider,

		wire.Bind(new(Repository.DivisiRepositoryHandler), new(*Repository.DivisiRepositoryImpl)),
		wire.Bind(new(Services.DivisiServiceHandler), new(*Services.DivisiServiceImpl)),
		wire.Bind(new(Controller.DivisiControllerHandler), new(*Controller.DivisiControllerImpl)),
	)))
	return &Controller.DivisiControllerImpl{}
}
