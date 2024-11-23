package Controller

import (
	"net/http"
	"user/Constant"
	"user/Model/Web"
	"user/Model/Web/Response"
	"user/Services"

	"github.com/gofiber/fiber/v2"
)

type DivisiControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type DivisiControllerImpl struct {
	service Services.DivisiServiceHandler
}

func DivisiControllerProvider(service Services.DivisiServiceHandler) *DivisiControllerImpl {
	return &DivisiControllerImpl{
		service: service,
	}
}

func (h *DivisiControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.DivisiCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create Divisi
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Divisi created successfully", nil))
}

func (h *DivisiControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.DivisiUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update Divisi
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Divisi updated successfully", nil))
}

func (h *DivisiControllerImpl) Delete(c *fiber.Ctx) error {
	// Get divisi ID from URL query
	divisiId, err := c.ParamsInt("divisiId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Divisi ID", err))
	}

	// Call service to delete Divisi
	if serviceErr := h.service.Delete(int64(divisiId)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Divisi deleted successfully", nil))
}

func (h *DivisiControllerImpl) FindById(c *fiber.Ctx) error {
	// Get divisi ID from URL query
	divisiId, err := c.ParamsInt("divisiId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Divisi ID", err))
	}

	// Call service to find Divisi by ID
	response, serviceErr := h.service.FindById(int64(divisiId))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Divisi found", response))
}

func (h *DivisiControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all Divisi
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All Divisi found", response))
}
