package Controller

import (
	"net/http"
	"user/Constant"
	"user/Model/Web"
	"user/Model/Web/Response"
	"user/Services"

	"github.com/gofiber/fiber/v2"
)

type JabatanControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type JabatanControllerImpl struct {
	service Services.JabatanServiceHandler
}

func JabatanControllerProvider(service Services.JabatanServiceHandler) *JabatanControllerImpl {
	return &JabatanControllerImpl{
		service: service,
	}
}

func (h *JabatanControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.JabatanCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create Jabatan
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Jabatan created successfully", nil))
}

func (h *JabatanControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.JabatanUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update Jabatan
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Jabatan updated successfully", nil))
}

func (h *JabatanControllerImpl) Delete(c *fiber.Ctx) error {
	// Get jabatan ID from URL params
	jabatanId, err := c.ParamsInt("jabatanId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Jabatan ID", err))
	}

	// Call service to delete Jabatan
	if serviceErr := h.service.Delete(int64(jabatanId)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Jabatan deleted successfully", nil))
}

func (h *JabatanControllerImpl) FindById(c *fiber.Ctx) error {
	// Get jabatan ID from URL params
	jabatanId, err := c.ParamsInt("jabatanId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Jabatan ID", err))
	}

	// Call service to find Jabatan by ID
	response, serviceErr := h.service.FindById(int64(jabatanId))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Jabatan found", response))
}

func (h *JabatanControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all Jabatan
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All Jabatan found", response))
}
