//go:build wireinject
// +build wireinject

package Route

import (
	"itam/Controller"
	"itam/Repository"
	"itam/Services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func assetDI(db *gorm.DB) *Controller.AssetControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.AssetRepositoryProvider,
		Services.AssetServiceControllerProvider,
		Controller.AssetControllerProvider,

		wire.Bind(new(Repository.AssetRepositoryHandler), new(*Repository.AssetRepositoryImpl)),
		wire.Bind(new(Services.AssetServiceHandler), new(*Services.AssetServiceImpl)),
		wire.Bind(new(Controller.AssetControllerHandler), new(*Controller.AssetControllerImpl)),
	),
	))
	return &Controller.AssetControllerImpl{}

}x

// Asset Perangkat DI
func assetPerangkatDI(db *gorm.DB) *Controller.AssetPerangkatControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.AssetPerangkatRepositoryProvider,
		Repository.AssetRepositoryProvider,
		Services.AssetPerangkatServiceProvider,
		Controller.AssetPerangkatControllerProvider,

		wire.Bind(new(Repository.AssetPerangkatRepositoryHandler), new(*Repository.AssetPerangkatRepositoryImpl)),
		wire.Bind(new(Repository.AssetRepositoryHandler), new(*Repository.AssetRepositoryImpl)),
		wire.Bind(new(Services.AssetPerangkatServiceHandler), new(*Services.AssetPerangkatServiceImpl)),
		wire.Bind(new(Controller.AssetPerangkatControllerHandler), new(*Controller.AssetPerangkatControllerImpl)),
	)))
	return &Controller.AssetPerangkatControllerImpl{}
}

// Detail Asset Aplikasi DI
func AssetAplikasiDI(db *gorm.DB) *Controller.AssetAplikasiControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.AssetAplikasiRepositoryProvider,
		Repository.AssetRepositoryProvider,
		Services.AssetAplikasiServiceProvider,
		Controller.AssetAplikasiControllerProvider,

		wire.Bind(new(Repository.AssetAplikasiRepositoryHandler), new(*Repository.AssetAplikasiRepositoryImpl)),
		wire.Bind(new(Repository.AssetRepositoryHandler), new(*Repository.AssetRepositoryImpl)),
		wire.Bind(new(Services.AssetAplikasiServiceHandler), new(*Services.AssetAplikasiServiceImpl)),
		wire.Bind(new(Controller.AssetAplikasiControllerHandler), new(*Controller.AssetAplikasiControllerImpl)),
	)))
	return &Controller.AssetAplikasiControllerImpl{}
}

// Asset Hardware DI
func assetHardwareDI(db *gorm.DB) *Controller.AssetHardwareControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.AssetHardwareRepositoryProvider,
		Repository.AssetRepositoryProvider,
		Services.AssetHardwareServiceProvider,
		Controller.AssetHardwareControllerProvider,

		wire.Bind(new(Repository.AssetHardwareRepositoryHandler), new(*Repository.AssetHardwareRepositoryImpl)),
		wire.Bind(new(Repository.AssetRepositoryHandler), new(*Repository.AssetRepositoryImpl)),
		wire.Bind(new(Services.AssetHardwareServiceHandler), new(*Services.AssetHardwareServiceImpl)),
		wire.Bind(new(Controller.AssetHardwareControllerHandler), new(*Controller.AssetHardwareControllerImpl)),
	)))
	return &Controller.AssetHardwareControllerImpl{}
}

// Detail Asset Lisensi DI
func AssetLisensiDI(db *gorm.DB) *Controller.AssetLisensiControllerImpl {
	panic(wire.Build(wire.NewSet(
		Repository.AssetLisensiRepositoryProvider,
		Repository.AssetRepositoryProvider,
		Services.AssetLisensiServiceProvider,
		Controller.AssetLisensiControllerProvider,

		wire.Bind(new(Repository.AssetLisensiRepositoryHandler), new(*Repository.AssetLisensiRepositoryImpl)),
		wire.Bind(new(Repository.AssetRepositoryHandler), new(*Repository.AssetRepositoryImpl)),
		wire.Bind(new(Services.AssetLisensiServiceHandler), new(*Services.AssetLisensiServiceImpl)),
		wire.Bind(new(Controller.AssetLisensiControllerHandler), new(*Controller.AssetLisensiControllerImpl)),
	)))
	return &Controller.AssetLisensiControllerImpl{}
}

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
