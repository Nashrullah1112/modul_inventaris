package Controller

import (
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type VendorControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type VendorControllerImpl struct {
	service Services.VendorServiceHandler
}

func VendorControllerProvider(service Services.VendorServiceHandler) *VendorControllerImpl {
	return &VendorControllerImpl{
		service: service,
	}
}

func (h *VendorControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.VendorCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create a new vendor
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusCreated).JSON(Web.SuccessResponse("Vendor created successfully", nil))
}

func (h *VendorControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.VendorUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)
	vendorId, err := c.ParamsInt("vendorId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.InvalidDataRequest, err))
	}
	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update an existing vendor
	if _, serviceErr = h.service.Update(int64(vendorId), request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Vendor updated successfully", nil))
}

func (h *VendorControllerImpl) Delete(c *fiber.Ctx) error {
	// Get vendor ID from URL params
	vendorId, err := c.ParamsInt("vendorId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.InvalidDataRequest, err))
	}

	// Call service to delete the vendor
	if serviceErr := h.service.Delete(int64(vendorId)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Vendor deleted successfully", nil))
}

func (h *VendorControllerImpl) FindById(c *fiber.Ctx) error {
	// Get vendor ID from URL params
	vendorId, err := c.ParamsInt("vendorId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.InvalidDataRequest, err))
	}

	// Call service to find the vendor by ID
	response, serviceErr := h.service.FindById(int64(vendorId))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Vendor found", response))
}

func (h *VendorControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all vendors
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All vendors found", response))
}
