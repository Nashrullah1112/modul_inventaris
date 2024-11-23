package Controller

import (
	"net/http"
	"user/Constant"
	"user/Model/Web"
	"user/Model/Web/Response"
	"user/Services"

	"github.com/gofiber/fiber/v2"
)

type UserControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	service Services.UserServiceHandler
}

func UserControllerProvider(service Services.UserServiceHandler) *UserControllerImpl {
	return &UserControllerImpl{
		service: service,
	}
}

func (h *UserControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.UserCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create a new user
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusCreated).JSON(Web.SuccessResponse("User created successfully", nil))
}

func (h *UserControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.UserUpdateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to update an existing user
	if _, serviceErr = h.service.Update(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("User updated successfully", nil))
}

func (h *UserControllerImpl) Delete(c *fiber.Ctx) error {
	// Get user ID from URL params
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid user ID", err))
	}

	// Call service to delete the user
	if serviceErr := h.service.Delete(int64(userId)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("User deleted successfully", nil))
}

func (h *UserControllerImpl) FindById(c *fiber.Ctx) error {
	// Get user ID from URL params
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid user ID", err))
	}

	// Call service to find the user by ID
	response, serviceErr := h.service.FindById(int64(userId))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("User found", response))
}

func (h *UserControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all users
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All users found", response))
}
