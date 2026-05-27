package handlers

import (
	"net/http"
	"order-service/models"
	"order-service/usecases"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(e echo.Context) error {
	var order models.OrderRequest

	if err := e.Bind(&order); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body", "details": err.Error(),
		})
	}

	if order.UserID.IsZero() || order.ProductID.IsZero() || order.Quantity <= 0 {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "User ID, Product ID and quantity are required and must be valid",
		})
	}

	err := usecases.CreateOrder(order)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create order", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, echo.Map{
		"message": "Order created successfully",
	})
}

func GetAllOrders(e echo.Context) error {
	orders, err := usecases.GetAllOrders()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to retrieve orders", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"orders": orders,
	})
}

func GetOrderByID(e echo.Context) error {
	idParam := e.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid order ID", "details": err.Error(),
		})
	}

	order, err := usecases.GetOrderByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{
			"error": "Order not found", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, order)
}

func UpdateOrder(e echo.Context) error {
	idParam := e.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid order ID", "details": err.Error(),
		})
	}

	var updateData models.OrderRequest
	if err := e.Bind(&updateData); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body", "details": err.Error(),
		})
	}

	if updateData.UserID.IsZero() || updateData.ProductID.IsZero() || updateData.Quantity <= 0 {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "User ID, Product ID and quantity are required and must be valid",
		})
	}

	err = usecases.UpdateOrder(id, updateData)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to update order", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message": "Order updated successfully",
	})
}

func DeleteOrder(e echo.Context) error {
	idParam := e.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid order ID", "details": err.Error(),
		})
	}

	err = usecases.DeleteOrder(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to delete order", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message": "Order deleted successfully",
	})
}

func UpdateOrderStatusCronHandler(c echo.Context) error {
	matchedCount, modifiedCount, err := usecases.UpdateOrderStatusCron()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update order status",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "order status updated successfully",
		"matched_count":  matchedCount,
		"modified_count": modifiedCount,
	})
}
