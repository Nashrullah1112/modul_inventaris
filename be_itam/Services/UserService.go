package Services

import (
	"fmt"
	"itam/Middleware"
	"itam/Model/Database"
	"itam/Model/Domain"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
	"log"
	"net/http"
	"os"
	"time"
)

type (
	UserServiceHandler interface {
		Create(request Response.UserCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Update(request Response.UserUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto)
		Delete(userId int64) (serviceErr *Web.ServiceErrorDto)
		FindById(userId int64) (user Response.UserResponse, serviceErr *Web.ServiceErrorDto)
		FindAll() (users []Response.UserResponse, serviceErr *Web.ServiceErrorDto)
		Login(request Domain.LoginRequest) (token Domain.JwtTokenDetail, serviceErr *Web.ServiceErrorDto)
		CheckRole(userId int64) (roleResponse Response.RoleResponse, serviceErr *Web.ServiceErrorDto)
		TotalUser() (total int64, serviceErr *Web.ServiceErrorDto)
		Seed() (seededUsers []Response.UserResponse, serviceErr *Web.ServiceErrorDto)
	}

	UserServiceImpl struct {
		userRepo    Repository.UserRepositoryHandler
		jabatanRepo Repository.JabatanRepositoryHandler
		divisiRepo  Repository.DivisiRepositoryHandler
		roleRepo    Repository.RoleRepositoryHandler
	}
)

func UserServiceProvider(userRepo Repository.UserRepositoryHandler, jabatanRepo Repository.JabatanRepositoryHandler, divisiRepo Repository.DivisiRepositoryHandler, roleRepo Repository.RoleRepositoryHandler) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo:    userRepo,
		jabatanRepo: jabatanRepo,
		divisiRepo:  divisiRepo,
		roleRepo:    roleRepo,
	}
}

func (h *UserServiceImpl) Seed() (seededUsers []Response.UserResponse, serviceErr *Web.ServiceErrorDto) {
	// Sample seed data
	usersToSeed := []Response.UserCreateRequest{
		{
			NIP:       12345678,
			Email:     "admin@gmail.com",
			JabatanID: 1,
			DivisiID:  1,
			RoleID:    1, // Role ID for the user
			Password:  "password",
			Nama:      "Administrator",
		},
		{
			NIP:       87654321,
			Email:     "user@example.com",
			JabatanID: 2,
			DivisiID:  2,
			RoleID:    2, // Role ID for the user
			Password:  "password",
			Nama:      "Regular User",
		},
	}

	// Sample Jabatan, Divisi, and Role data
	jabatanToSeed := []Response.JabatanCreateRequest{
		{Nama: "Admin"},
		{Nama: "User"},
	}
	divisiToSeed := []Response.DivisiCreateRequest{
		{Nama: "IT"},
		{Nama: "HR"},
	}
	roleToSeed := []Response.RoleCreateRequest{
		{Nama: "Admin"},
		{Nama: "User"},
	}
	moduleToSeed := []Response.ModuleCreateRequest{
		{Name: "User"},
		{Name: "Aplikasi"},
		{Name: "Perangkat"},
		{Name: "Lisensi"},
		{Name: "Hardware"},
		{Name: "Asset"},
		{Name: "Divisi"},
		{Name: "Jabatan"},
		{Name: "Role"},
		{Name: "Dashboard"},
		{Name: "Approval"},
		{Name: "Disposal"},
	}

	// Seed Jabatan if they don't exist and update the ID map
	jabatanIDMap := make(map[string]int64)
	for _, jabatan := range jabatanToSeed {
		existingJabatan, err := h.jabatanRepo.FindByNama(jabatan.Nama)
		if err != nil {
			// Create if not exists
			log.Printf("here")
			createdJabatan, err := h.jabatanRepo.Save(&Database.Jabatan{Nama: jabatan.Nama})
			if err != nil {
				log.Printf("Failed to seed jabatan: %s, error: %v\n", jabatan.Nama, err)
				return nil, Web.NewInternalServiceError(err)
			}
			jabatanIDMap[jabatan.Nama] = createdJabatan
		} else {
			jabatanIDMap[jabatan.Nama] = int64(existingJabatan.ID)
		}
	}
	log.Printf("done")
	log.Println(jabatanIDMap)
	// Seed Divisi if they don't exist and update the ID map
	divisiIDMap := make(map[string]int64)
	for _, divisi := range divisiToSeed {
		existingDivisi, err := h.divisiRepo.FindByNama(divisi.Nama)
		if err != nil {
			// Create if not exists
			log.Printf("here")
			createdDivisi, err := h.divisiRepo.Save(&Database.Divisi{Nama: divisi.Nama})
			if err != nil {
				log.Printf("Failed to seed divisi: %s, error: %v\n", divisi.Nama, err)
				return nil, Web.NewInternalServiceError(err)
			}
			divisiIDMap[divisi.Nama] = createdDivisi
		} else {
			divisiIDMap[divisi.Nama] = existingDivisi.ID
		}
	}
	log.Printf("done")
	log.Println(divisiIDMap)
	// Seed Role if they don't exist and update the ID map
	moduleIDMap := make(map[string]int64)
	for _, module := range moduleToSeed {
		existingModule, err := h.roleRepo.FindModuleByNama(module.Name)
		if err != nil {
			createdModule, err := h.roleRepo.SaveModule(&Database.Module{Nama: module.Name})
			if err != nil {
				log.Printf("Failed to seed module: %s, error: %v\n", module.Name, err)
				return nil, Web.NewInternalServiceError(err)
			}
			moduleIDMap[module.Name] = createdModule
		} else {
			moduleIDMap[module.Name] = existingModule.ID
		}
	}
	roleIDMap := make(map[string]int64)
	for _, role := range roleToSeed {
		log.Printf("here")

		existingRole, err := h.roleRepo.FindRoleByNama(role.Nama)
		if err != nil {
			// Create if not exists
			log.Println(err)
			log.Printf("here")

			createdRole, err := h.roleRepo.SaveRole(&Database.Role{Nama: role.Nama})
			if err != nil {
				log.Printf("Failed to seed role: %s, error: %v\n", role.Nama, err)
				return nil, Web.NewInternalServiceError(err)
			}
			roleIDMap[role.Nama] = createdRole
			modules, err := h.roleRepo.FindAllModules()
			if err != nil {
				log.Printf("Failed to seed role: %s, error: %v\n", role.Nama, err)
				return nil, Web.NewInternalServiceError(err)
			}
			for _, module := range modules {
				check, err := h.roleRepo.CheckRoleModuleExist(existingRole.ID, module.ID)
				if !check && err != nil {
					continue
				}
				_, err = h.roleRepo.SaveRoleModule(&Database.RoleModule{
					RoleID:    existingRole.ID,
					ModuleID:  module.ID,
					IsAllowed: true,
				})
				if err != nil {
					log.Printf("Failed to seed role module: %s, error: %v\n", role.Nama, err)
					return nil, Web.NewInternalServiceError(err)
				}
			}
		} else {
			roleIDMap[role.Nama] = existingRole.ID
			modules, err := h.roleRepo.FindAllModules()
			if err != nil {
				log.Printf("Failed to seed role: %s, error: %v\n", role.Nama, err)
				return nil, Web.NewInternalServiceError(err)
			}
			for _, module := range modules {
				check, err := h.roleRepo.CheckRoleModuleExist(existingRole.ID, module.ID)
				if !check && err != nil {
					continue
				}
				_, err = h.roleRepo.SaveRoleModule(&Database.RoleModule{
					RoleID:    existingRole.ID,
					ModuleID:  module.ID,
					IsAllowed: true,
				})
				if err != nil {
					log.Printf("Failed to seed role module: %s, error: %v\n", role.Nama, err)
					return nil, Web.NewInternalServiceError(err)
				}
			}
		}
	}
	log.Printf("done")
	log.Println(roleIDMap)

	// Iterate over seed data for users
	for _, user := range usersToSeed {
		// Use the appropriate IDs from the maps (jabatanIDMap, divisiIDMap, roleIDMap)
		user.JabatanID = jabatanIDMap["Admin"] // Example: using "Admin" for JabatanID
		user.DivisiID = divisiIDMap["IT"]      // Example: using "IT" for DivisiID
		user.RoleID = roleIDMap["Admin"]       // Example: using "Admin" for RoleID
		log.Printf("here")
		existingUser, err := h.userRepo.FindByEmail(user.Email)
		if err == nil && existingUser.ID != 0 {
			_, err = h.userRepo.Update(&Database.User{
				NIP:              user.NIP,
				Email:            user.Email,
				JabatanID:        user.JabatanID,
				DivisiID:         user.DivisiID,
				RoleID:           user.RoleID,
				Password:         user.Password,
				Nama:             user.Nama,
				TanggalBergabung: time.Now(),
			})
			if err != nil {
				log.Printf("Failed to seed user: %s, error: %v\n", user.Email, err)
				return nil, Web.NewInternalServiceError(err)
			}
		}

		id, err := h.userRepo.Save(&Database.User{
			NIP:              user.NIP,
			Email:            user.Email,
			JabatanID:        user.JabatanID,
			DivisiID:         user.DivisiID,
			RoleID:           user.RoleID,
			Password:         user.Password,
			Nama:             user.Nama,
			TanggalBergabung: time.Now(),
		})
		if err != nil {
			log.Printf("Failed to seed user: %s, error: %v\n", user.Email, err)
			return nil, Web.NewInternalServiceError(err)
		}

		createdUser, err := h.userRepo.FindById(id)
		if err == nil {
			jabatan, _ := h.jabatanRepo.FindById(createdUser.JabatanID)
			divisi, _ := h.divisiRepo.FindById(createdUser.DivisiID)

			seededUsers = append(seededUsers, Response.UserResponse{
				ID:        createdUser.ID,
				NIP:       createdUser.NIP,
				Nama:      createdUser.Nama,
				Email:     createdUser.Email,
				JabatanID: createdUser.JabatanID,
				Jabatan: Response.JabatanResponse{
					ID:   jabatan.ID,
					Nama: jabatan.Nama,
				},
				DivisiID: createdUser.DivisiID,
				Divisi: Response.DivisiResponse{
					ID:   divisi.ID,
					Nama: divisi.Nama,
				},
				RoleID: createdUser.RoleID,

				TanggalBergabung: createdUser.TanggalBergabung,
			})
		}
	}

	log.Println("User seeding completed successfully.")
	return seededUsers, nil
}
func (h *UserServiceImpl) Create(request Response.UserCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	user, _ := h.userRepo.FindByEmail(request.Email)
	if user.ID != 0 {
		return 0, Web.NewCustomServiceError("Email already exist", nil, http.StatusConflict)
	}

	id, err := h.userRepo.Save(&Database.User{
		NIP:              request.NIP,
		Email:            request.Email,
		JabatanID:        request.JabatanID,
		DivisiID:         request.DivisiID,
		RoleID:           request.RoleID,
		Password:         request.Password,
		Nama:             request.Nama,
		TanggalBergabung: time.Now(),
	})
	if err != nil {
		return 0, Web.NewCustomServiceError("User not created", err, http.StatusInternalServerError)
	}

	return id, nil
}

func (h *UserServiceImpl) Update(request Response.UserUpdateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	existingUser, err := h.userRepo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}

	id, err = h.userRepo.Update(&Database.User{
		ID:               existingUser.ID,
		NIP:              request.NIP,
		Email:            request.Email,
		JabatanID:        request.JabatanID,
		DivisiID:         request.DivisiID,
		RoleID:           request.RoleID,
		TanggalBergabung: existingUser.TanggalBergabung,
	})
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *UserServiceImpl) Delete(userId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.userRepo.FindById(userId)
	if err != nil {
		return Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}

	if err := h.userRepo.Delete(userId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *UserServiceImpl) FindById(userId int64) (user Response.UserResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.userRepo.FindById(userId)
	if err != nil {
		return Response.UserResponse{}, Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}
	jabatan, err := h.jabatanRepo.FindById(data.JabatanID)
	if err != nil {
		return Response.UserResponse{}, Web.NewCustomServiceError("Jabatan not found", err, http.StatusNotFound)
	}
	divisi, err := h.divisiRepo.FindById(data.DivisiID)
	if err != nil {
		return Response.UserResponse{}, Web.NewCustomServiceError("Divisi not found", err, http.StatusNotFound)
	}
	user = Response.UserResponse{
		ID:        data.ID,
		NIP:       data.NIP,
		Nama:      data.Nama,
		Email:     data.Email,
		JabatanID: data.JabatanID,
		Jabatan: Response.JabatanResponse{
			ID:   jabatan.ID,
			Nama: jabatan.Nama,
		},
		DivisiID: data.DivisiID,
		Divisi: Response.DivisiResponse{
			ID:   divisi.ID,
			Nama: divisi.Nama,
		},
		TanggalBergabung: data.TanggalBergabung,
	}

	return user, nil
}

func (h *UserServiceImpl) FindAll() (users []Response.UserResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.userRepo.FindAll()
	if err != nil {
		return []Response.UserResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		jabatan, err := h.jabatanRepo.FindById(d.JabatanID)
		if err != nil {
			continue
		}
		divisi, err := h.divisiRepo.FindById(d.DivisiID)
		if err != nil {
			continue
		}
		users = append(users, Response.UserResponse{
			ID:        d.ID,
			NIP:       d.NIP,
			Nama:      d.Nama,
			Email:     d.Email,
			JabatanID: d.JabatanID,
			Jabatan: Response.JabatanResponse{
				ID:   jabatan.ID,
				Nama: jabatan.Nama,
			},
			DivisiID: d.DivisiID,
			Divisi: Response.DivisiResponse{
				ID:   divisi.ID,
				Nama: divisi.Nama,
			},
			TanggalBergabung: d.TanggalBergabung,
		})
	}

	return users, nil
}

func (h *UserServiceImpl) Login(request Domain.LoginRequest) (token Domain.JwtTokenDetail, serviceErr *Web.ServiceErrorDto) {
	log.Println(request)
	data, err := h.userRepo.FindByEmail(request.Email)
	log.Println(data)
	if err != nil {
		return Domain.JwtTokenDetail{}, Web.NewCustomServiceError("User dan Password salah", err, http.StatusUnauthorized)
	}

	if !Middleware.CheckPasswordHash(request.Password, []byte(data.Password)) {
		return Domain.JwtTokenDetail{}, Web.NewCustomServiceError("User dan Password salah", nil, http.StatusUnauthorized)
	}

	tokenPtr, err := Middleware.GenerateTokenJwtV2(time.Hour*24*7, data.ID, os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))
	if err != nil {
		return Domain.JwtTokenDetail{}, Web.NewCustomServiceError("Login failed", err, http.StatusUnauthorized)
	}
	token = *tokenPtr

	return token, nil
}

func (h *UserServiceImpl) CheckRole(userId int64) (roleResponse Response.RoleResponse, serviceErr *Web.ServiceErrorDto) {
	user, err := h.userRepo.FindById(userId)
	if err != nil {
		return Response.RoleResponse{}, Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}
	fmt.Println(user.JabatanID)
	role, err := h.roleRepo.FindRoleById(user.RoleID)
	if err != nil {
		return Response.RoleResponse{}, Web.NewCustomServiceError("Jabatan not found", err, http.StatusNotFound)
	}
	module, err := h.roleRepo.FindAllModulesByRole(role.ID)
	if err != nil {
		return Response.RoleResponse{}, Web.NewCustomServiceError("Module not found", err, http.StatusNotFound)
	}
	return Response.RoleResponse{
		ID:      role.ID,
		Name:    role.Nama,
		Modules: module,
	}, nil
}
func (h *UserServiceImpl) TotalUser() (total int64, serviceErr *Web.ServiceErrorDto) {
	total, err := h.userRepo.TotalUser()
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}
	return total, nil
}
