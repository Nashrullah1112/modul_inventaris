package Controller

import (
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AssetControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type AssetControllerImpl struct {
	service Services.AssetServiceHandler
}

func AssetControllerProvider(service Services.AssetServiceHandler) *AssetControllerImpl {
	return &AssetControllerImpl{
		service: service,
	}
}

func (h *AssetControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.AssetCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create asset
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Asset created successfully", nil))
}

func (h *AssetControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.AssetUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update asset
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Asset updated successfully", nil))
}

func (h *AssetControllerImpl) Delete(c *fiber.Ctx) error {
	// Get asset ID from URL query
	assetID, err := c.ParamsInt("assetId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid asset ID", err))
	}

	// Call service to delete asset
	if serviceErr := h.service.Delete(int64(assetID)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Asset deleted successfully", nil))
}

func (h *AssetControllerImpl) FindById(c *fiber.Ctx) error {
	// Get asset ID from URL query
	assetID, err := c.ParamsInt("assetId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid asset ID", err))
	}

	// Call service to find asset by ID
	response, serviceErr := h.service.FindById(int64(assetID))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Asset found", response))
}

func (h *AssetControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all assets
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All assets found", response))
}
