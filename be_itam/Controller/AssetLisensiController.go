package Controller

import (
	"itam/Constant"
	"itam/Model/Web"
	"itam/Model/Web/Response"
	"itam/Services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AssetLisensiControllerHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	TotalLisensi(c *fiber.Ctx) error
}

type AssetLisensiControllerImpl struct {
	service Services.AssetLisensiServiceHandler
}

func AssetLisensiControllerProvider(service Services.AssetLisensiServiceHandler) *AssetLisensiControllerImpl {
	return &AssetLisensiControllerImpl{
		service: service,
	}
}

// Create DetailAsetLisensi
func (h *AssetLisensiControllerImpl) Create(c *fiber.Ctx) error {
	var (
		request    Response.AssetLicenseCreateRequest
		serviceErr *Web.ServiceErrorDto
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}

	// Call service to create DetailAsetLisensi
	if _, serviceErr = h.service.Create(request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetailAsetLisensi created successfully", nil))
}

// Update DetailAsetLisensi
func (h *AssetLisensiControllerImpl) Update(c *fiber.Ctx) error {
	var (
		request    Response.AssetLicenseUpdateRequest
		serviceErr *Web.ServiceErrorDto
		lisensiID  int
	)

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse(Constant.FailedBindError, nil))
	}
	lisensiID, err := c.ParamsInt("lisensiID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid Aplikasi ID", err))
	}
	// Call service to update DetailAsetLisensi
	if _, serviceErr = h.service.Update(int64(lisensiID), request); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetailAsetLisensi updated successfully", nil))
}

// Delete DetailAsetLisensi
func (h *AssetLisensiControllerImpl) Delete(c *fiber.Ctx) error {
	// Get DetailAsetLisensi ID from URL query
	detailAsetlisensiID, err := c.ParamsInt("lisensiID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid DetailAsetLisensi ID", err))
	}

	// Call service to delete DetailAsetLisensi
	if serviceErr := h.service.Delete(int64(detailAsetlisensiID)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetailAsetLisensi deleted successfully", nil))
}

// FindById DetailAsetLisensi
func (h *AssetLisensiControllerImpl) FindById(c *fiber.Ctx) error {
	// Get DetailAsetLisensi ID from URL query
	detailAsetlisensiID, err := c.ParamsInt("lisensiID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid DetailAsetLisensi ID", err))
	}

	// Call service to find DetailAsetLisensi by ID
	response, serviceErr := h.service.FindById(int64(detailAsetlisensiID))
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("DetailAsetLisensi found", response))
}

// FindAll DetailAsetLisensi
func (h *AssetLisensiControllerImpl) FindAll(c *fiber.Ctx) error {
	// Call service to find all DetailAsetLisensis
	response, serviceErr := h.service.FindAll()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All DetailAsetLisensis found", response))
}

// TotalLisensi returns the total number of DetailAsetLisensis
func (h *AssetLisensiControllerImpl) TotalLisensi(c *fiber.Ctx) error {
	// Call service to get the total number of DetailAsetLisensis
	total, serviceErr := h.service.TotalLisensi()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Total DetailAsetLisensis retrieved successfully", total))
}

func (h *AssetLisensiControllerImpl) GetNotifications(c *fiber.Ctx) error {
	// Call service to get all notifications
	notifications, serviceErr := h.service.NotifyExpiringLicenses()
	if serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("All notifications found", notifications))
}

func (h *AssetLisensiControllerImpl) MarkNotificationAsRead(c *fiber.Ctx) error {
	// Get notification ID from URL query
	notificationID, err := c.ParamsInt("notificationID")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Web.ErrorResponse("Invalid notification ID", err))
	}

	// Call service to mark notification as read
	if serviceErr := h.service.MarkNotificationAsRead(int64(notificationID)); serviceErr != nil {
		return c.Status(serviceErr.StatusCode).JSON(Web.ErrorResponse(serviceErr.Message, serviceErr.Err))
	}

	return c.Status(http.StatusOK).JSON(Web.SuccessResponse("Notification marked as read successfully", nil))
}
