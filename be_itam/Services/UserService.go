package Services

import (
	"itam/Middleware"
	"itam/Model/Database"
	"itam/Model/Domain"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Repository"
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
	}

	UserServiceImpl struct {
		repo Repository.UserRepositoryHandler
	}
)

func UserServiceProvider(repo Repository.UserRepositoryHandler) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (h *UserServiceImpl) Create(request Response.UserCreateRequest) (id int64, serviceErr *Web.ServiceErrorDto) {
	user, _ := h.repo.FindByEmail(request.Email)
	if user.ID != 0 {
		return 0, Web.NewCustomServiceError("Email already exist", nil, http.StatusConflict)
	}

	id, err := h.repo.Save(&Database.User{
		NIP:              request.NIP,
		Email:            request.Email,
		JabatanID:        request.JabatanID,
		DivisiID:         request.DivisiID,
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
	existingUser, err := h.repo.FindById(request.ID)
	if err != nil {
		return 0, Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}

	id, err = h.repo.Update(&Database.User{
		ID:               existingUser.ID,
		NIP:              request.NIP,
		Email:            request.Email,
		JabatanID:        request.JabatanID,
		DivisiID:         request.DivisiID,
		TanggalBergabung: existingUser.TanggalBergabung,
	})
	if err != nil {
		return 0, Web.NewInternalServiceError(err)
	}

	return id, nil
}

func (h *UserServiceImpl) Delete(userId int64) (serviceErr *Web.ServiceErrorDto) {
	_, err := h.repo.FindById(userId)
	if err != nil {
		return Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}

	if err := h.repo.Delete(userId); err != nil {
		return Web.NewInternalServiceError(err)
	}

	return nil
}

func (h *UserServiceImpl) FindById(userId int64) (user Response.UserResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindById(userId)
	if err != nil {
		return Response.UserResponse{}, Web.NewCustomServiceError("User not found", err, http.StatusNotFound)
	}

	user = Response.UserResponse{
		ID:               data.ID,
		NIP:              data.NIP,
		Email:            data.Email,
		JabatanID:        data.JabatanID,
		DivisiID:         data.DivisiID,
		TanggalBergabung: data.TanggalBergabung,
	}

	return user, nil
}

func (h *UserServiceImpl) FindAll() (users []Response.UserResponse, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindAll()
	if err != nil {
		return []Response.UserResponse{}, Web.NewInternalServiceError(err)
	}

	for _, d := range data {
		users = append(users, Response.UserResponse{
			ID:               d.ID,
			NIP:              d.NIP,
			Email:            d.Email,
			JabatanID:        d.JabatanID,
			DivisiID:         d.DivisiID,
			TanggalBergabung: d.TanggalBergabung,
		})
	}

	return users, nil
}

func (h *UserServiceImpl) Login(request Domain.LoginRequest) (token Domain.JwtTokenDetail, serviceErr *Web.ServiceErrorDto) {
	data, err := h.repo.FindByEmail(request.Email)
	if err != nil {
		return Domain.JwtTokenDetail{}, Web.NewCustomServiceError("User dan Password salah", err, http.StatusNotFound)
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
