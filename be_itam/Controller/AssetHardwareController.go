package Controller

import (
	"fmt"
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AssetHardwareControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FormAssetHardware(c *fiber.Ctx) error
	TotalHardware(c *fiber.Ctx) error
}

type AssetHardwareControllerImpl struct {
	service Services.AssetHardwareServiceHandler
}

func AssetHardwareControllerProvider(service Services.AssetHardwareServiceHandler) *AssetHardwareControllerImpl {
	return &AssetHardwareControllerImpl{
		service: service,
	}
}

func (h *AssetHardwareControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.DetailAsetHardwareCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create detail aset hardware
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail asset hardware created successfully", nil))
}

func (h *AssetHardwareControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.AssetHardwareUpdateRequest
		serviceErr *Web.ServiceErrorDto
		hardwareID int
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}
	hardwareID, err := c.ParamsInt("hardwareID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid asset hardware ID", err))
	}
	fmt.Println(hardwareID)

	if _, serviceErr = h.service.UpdateAssetHardware(int64(hardwareID), request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail asset hardware updated successfully", nil))
}

func (h *AssetHardwareControllerImpl) Delete(c *fiber.Ctx) error {
	// Get detail asset hardware ID from URL params
	detailAsethardwareID, err := c.ParamsInt("hardwareID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid detail asset hardware ID", err))
	}

	// Call service to delete detail asset hardware
	if serviceErr := h.service.Delete(int64(detailAsethardwareID)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail asset hardware deleted successfully", nil))
}

func (h *AssetHardwareControllerImpl) FindById(c *fiber.Ctx) error {
	// Get detail asset hardware ID from URL params
	detailAsethardwareID, err := c.ParamsInt("hardwareID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid detail asset hardware ID", err))
	}

	// Call service to find detail asset hardware by ID
	response, serviceErr := h.service.FindById(int64(detailAsethardwareID))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail asset hardware found", response))
}

func (h *AssetHardwareControllerImpl) FindAll(c *fiber.Ctx) error {

	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All detail asset hardware found", response))
}
func (h *AssetHardwareControllerImpl) FormAssetHardware(c *fiber.Ctx) error {
	var (
		request    Response.AssetHardwareCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	file, err := c.FormFile("tanda_terima")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, err))
	}

	// Tentukan lokasi tujuan penyimpanan file
	destination := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), strings.ReplaceAll(file.Filename, " ", ""))

	// Simpan file ke folder yang ditentukan
	if err = c.SaveFile(file, "./public/static/hardware/tanda_terima/"+destination); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Web.ErrorResponse(Constant.InternalHttpError, err))
	}

	fileNota, err := c.FormFile("nota_pembelian")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, err))
	}

	// Tentukan lokasi tujuan penyimpanan file
	destinationNota := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), strings.ReplaceAll(fileNota.Filename, " ", ""))

	// Simpan file ke folder yang ditentukan
	if err = c.SaveFile(fileNota, "./public/static/hardware/nota_pembelian/"+destinationNota); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Web.ErrorResponse(Constant.InternalHttpError, err))
	}

	// Parse request body (selain file)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}
	log.Println(request)

	// Call service untuk menyimpan data asset hardware
	if _, serviceErr = h.service.FormAssetHardware(request, destination, destinationNota); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Detail asset hardware created successfully", nil))
}
func (h *AssetHardwareControllerImpl) TotalHardware(c *fiber.Ctx) error {
	total, serviceErr := h.service.TotalHardware()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Total asset hardware retrieved successfully", total))
}
