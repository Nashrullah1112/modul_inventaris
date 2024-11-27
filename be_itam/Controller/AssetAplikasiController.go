package Controller

import (
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AssetAplikasiControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type AssetAplikasiControllerImpl struct {
	service Services.AssetAplikasiServiceHandler
}

func AssetAplikasiControllerProvider(service Services.AssetAplikasiServiceHandler) *AssetAplikasiControllerImpl {
	return &AssetAplikasiControllerImpl{
		service: service,
	}
}

func (h *AssetAplikasiControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.DetaiAsetAplikasiCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create detaiAsetAplikasi
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetaiAsetAplikasi created successfully", nil))
}

func (h *AssetAplikasiControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.DetaiAsetAplikasiUpdateRequest
		serviceErr *Web.ServiceErrorDto
		aplikasiID int
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	aplikasiID, err := c.ParamsInt("aplikasiID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Aplikasi ID", err))
	}
	// Call service to update detaiAsetAplikasi
	if _, serviceErr = h.service.Update(int64(aplikasiID), request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetaiAsetAplikasi updated successfully", nil))
}

func (h *AssetAplikasiControllerImpl) Delete(c *fiber.Ctx) error {
	// Get detaiAsetAplikasi ID from URL query
	detaiAsetaplikasiID, err := c.ParamsInt("aplikasiID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid DetaiAsetAplikasi ID", err))
	}

	// Call service to delete detaiAsetAplikasi
	if serviceErr := h.service.Delete(int64(detaiAsetaplikasiID)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetaiAsetAplikasi deleted successfully", nil))
}

func (h *AssetAplikasiControllerImpl) FindById(c *fiber.Ctx) error {
	// Get detaiAsetAplikasi ID from URL query
	detaiAsetaplikasiID, err := c.ParamsInt("aplikasiID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid DetaiAsetAplikasi ID", err))
	}

	// Call service to find detaiAsetAplikasi by ID
	response, serviceErr := h.service.FindById(int64(detaiAsetaplikasiID))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetaiAsetAplikasi found", response))
}

func (h *AssetAplikasiControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all detaiAsetAplikasis
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All DetaiAsetAplikasis found", response))
}

func (h *AssetAplikasiControllerImpl) TotalAplikasi(c *fiber.Ctx) error {
	// Call service to get total count of aplikasi
	total, serviceErr := h.service.TotalAplikasi()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Total Aplikasi count retrieved", total))
}
