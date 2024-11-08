package Controller

import (
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RoleControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type RoleControllerImpl struct {
	service Services.RoleServiceHandler
}

func RoleControllerProvider(service Services.RoleServiceHandler) *RoleControllerImpl {
	return &RoleControllerImpl{
		service: service,
	}
}

func (h *RoleControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.RoleCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create a new role
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Role created successfully", nil))
}

func (h *RoleControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.RoleUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update an existing role
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Role updated successfully", nil))
}

func (h *RoleControllerImpl) Delete(c *fiber.Ctx) error {
	// Get role ID from URL params
	roleId, err := c.ParamsInt("roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid role ID", err))
	}

	// Call service to delete the role
	if serviceErr := h.service.Delete(int64(roleId)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Role deleted successfully", nil))
}

func (h *RoleControllerImpl) FindById(c *fiber.Ctx) error {
	// Get role ID from URL params
	roleId, err := c.ParamsInt("roleId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid role ID", err))
	}

	// Call service to find the role by ID
	response, serviceErr := h.service.FindById(int64(roleId))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Role found", response))
}

func (h *RoleControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all roles
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All roles found", response))
}
