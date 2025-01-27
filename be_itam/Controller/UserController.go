package Controller

import (
	"itam/Constant"
	"itam/Model/Domain"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	CheckRole(c *fiber.Ctx) error
	TotalUser(c *fiber.Ctx) error
	Seed(c *fiber.Ctx) error
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
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.InvalidDataRequest, err))
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
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.InvalidDataRequest, err))
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

func (h *UserControllerImpl) Login(c *fiber.Ctx) error {
	var (
		request    Domain.LoginRequest
		response   Domain.JwtTokenDetail
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to login
	response, serviceErr = h.service.Login(request)
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Login successful", response))
}
func (h *UserControllerImpl) CheckRole(c *fiber.Ctx) error {
	// Get user ID from local context
	userId, ok := c.Locals("userID").(int64)

	if !ok || userId == 0 {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.InvalidDataRequest, nil))
	}
	// Call service to check the role of the user
	response, serviceErr := h.service.CheckRole(userId)
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Role checked successfully", response))
}

func (h *UserControllerImpl) TotalUser(c *fiber.Ctx) error {
	total, serviceErr := h.service.TotalUser()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Total user found", total))
}
func (h *UserControllerImpl) Seed(c *fiber.Ctx) error {
	_, serviceErr := h.service.Seed()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("User seeded successfully", nil))
}
