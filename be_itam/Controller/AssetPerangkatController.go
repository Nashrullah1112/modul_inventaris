package Controller

import (
	"fmt"
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"
	"strings"
	"time"

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
	fileTerima, err := c.FormFile("tanda_terima")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, err))
	}

	// Tentukan lokasi tujuan penyimpanan fileTerima
	destinationTerima := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), strings.ReplaceAll(fileTerima.Filename, " ", ""))

	// Simpan fileTerima ke folder yang ditentukan
	if err = c.SaveFile(fileTerima, "./public/static/perangkat/tanda_terima/"+destinationTerima); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Web.ErrorResponse(Constant.InternalHttpError, err))
	}

	fileNota, err := c.FormFile("nota_pembelian")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, err))
	}

	// Tentukan lokasi tujuan penyimpanan file
	destinationNota := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), strings.ReplaceAll(fileNota.Filename, " ", ""))

	// Simpan file ke folder yang ditentukan
	if err = c.SaveFile(fileNota, "./public/static/perangkat/nota_pembelian/"+destinationNota); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Web.ErrorResponse(Constant.InternalHttpError, err))
	}

	filePemeriksaan, err := c.FormFile("hasil_pemeriksaan")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, err))
	}

	// Tentukan lokasi tujuan penyimpanan file
	destinationPemeriksaan := fmt.Sprintf("%s-%s", time.Now().Format("20060102"), strings.ReplaceAll(filePemeriksaan.Filename, " ", ""))

	// Simpan file ke folder yang ditentukan
	if err = c.SaveFile(fileNota, "./public/static/perangkat/hasil_pemeriksaan/"+destinationPemeriksaan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Web.ErrorResponse(Constant.InternalHttpError, err))
	}

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create DetailAsetPerangkat
	if _, serviceErr = h.service.Create(request, destinationTerima, destinationNota, destinationPemeriksaan); serviceErr != nil {
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
