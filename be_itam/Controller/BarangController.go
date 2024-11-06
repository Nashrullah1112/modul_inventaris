package Controller

import (
	"fmt"
	"itam/Constant"
	"itam/Model/Web"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BarangControllerHandler interface {
	Create(c *fiber.Ctx) (err error)
	Update(c *fiber.Ctx) (err error)
	Delete(c *fiber.Ctx) (err error)
	FindById(c *fiber.Ctx) (err error)
	FindAll(c *fiber.Ctx) (err error)
}
type BarangControllerImpl struct {
	service Services.BarangServiceHandler
}

func BarangControllerControllerProvider(service Services.BarangServiceHandler) *BarangControllerImpl {
	return &BarangControllerImpl{
		service: service,
	}
}
func (h *BarangControllerImpl) Create(c *fiber.Ctx) (err error) {
	var (
		request    []Web.BarangCreateRequest
		serviceErr *Web.ServiceErrorDto
	)
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		fmt.Println("error at create barang service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("create succeed", nil))
}
func (h *BarangControllerImpl) Update(c *fiber.Ctx) (err error) {
	var (
		request    Web.BarangUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err = c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update barang
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		fmt.Println("error at update barang service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("update succeed", nil))
}

func (h *BarangControllerImpl) Delete(c *fiber.Ctx) error {
	// Get barang ID from URL query
	var serviceErr *Web.ServiceErrorDto

	barangID, err := c.ParamsInt("barangId")

	// Convert barang ID to int

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("invalid barang ID", err))
	}

	// Call service to delete barang
	if serviceErr = h.service.Delete(int64(barangID)); serviceErr != nil {
		fmt.Println("error at delete barang service: ", serviceErr.Err)
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("delete succeed", nil))
}

func (h *BarangControllerImpl) FindById(c *fiber.Ctx) error {
	var (
		serviceErr *Web.ServiceErrorDto
		response   Web.BarangResponse
		id, _      = c.ParamsInt("barangId")
	)

	if response, serviceErr = h.service.FindById(int64(id)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data barang", response))
}
func (h *BarangControllerImpl) FindAll(c *fiber.Ctx) (err error) {
	var (
		serviceErr *Web.ServiceErrorDto
		response   []Web.BarangResponse
	)

	if response, serviceErr = h.service.FindAll(); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("data barang", response))
}
