package Controller

import (
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AssetPerangkatControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	TotalPerangkat(c *fiber.Ctx) error
}

type AssetPerangkatControllerImpl struct {
	service Services.AssetPerangkatServiceHandler
}

func AssetPerangkatControllerProvider(service Services.AssetPerangkatServiceHandler) *AssetPerangkatControllerImpl {
	return &AssetPerangkatControllerImpl{
		service: service,
	}
}

func (h *AssetPerangkatControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.DetailAsetPerangkatCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create DetailAsetPerangkat
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail Aset Perangkat created successfully", nil))
}

func (h *AssetPerangkatControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request     Response.DetailAsetPerangkatUpdateRequest
		serviceErr  *Web.ServiceErrorDto
		perangkatID int
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}
	perangkatID, err := c.ParamsInt("perangkatID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Aplikasi ID", err))
	}
	// Call service to update DetailAsetPerangkat
	if _, serviceErr = h.service.Update(int64(perangkatID), request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail Aset Perangkat updated successfully", nil))
}

func (h *AssetPerangkatControllerImpl) Delete(c *fiber.Ctx) error {
	// Get detailAsetPerangkat ID from URL query
	detailAsetperangkatID, err := c.ParamsInt("perangkatID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Detail Aset Perangkat ID", err))
	}

	// Call service to delete DetailAsetPerangkat
	if serviceErr := h.service.Delete(int64(detailAsetperangkatID)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail Aset Perangkat deleted successfully", nil))
}

func (h *AssetPerangkatControllerImpl) FindById(c *fiber.Ctx) error {
	// Get detailAsetPerangkat ID from URL query
	detailAsetperangkatID, err := c.ParamsInt("perangkatID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Detail Aset Perangkat ID", err))
	}

	// Call service to find DetailAsetPerangkat by ID
	response, serviceErr := h.service.FindById(int64(detailAsetperangkatID))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail Aset Perangkat found", response))
}

func (h *AssetPerangkatControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all detailAsetPerangkat
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All Detail Aset Perangkat found", response))
}

func (h *AssetPerangkatControllerImpl) TotalPerangkat(c *fiber.Ctx) error {

	// Call service to get total perangkat
	total, serviceErr := h.service.TotalPerangkat()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Total Perangkat found", total))
}
