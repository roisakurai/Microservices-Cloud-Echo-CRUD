package handlers

import (
	"net/http"
	"product-service/models"
	"product-service/usecases" // Import usecases

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(e echo.Context) error {
	var product models.Product

	if err := e.Bind(&product); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body", "details": err.Error(),
		})
	}

	if product.Name == "" || product.Price <= 0 || product.Stock < 0 {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Name, price, and stock are required and must be valid",
		})
	}

	err := usecases.CreateProduct(product)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create product", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, echo.Map{
		"message": "Product created successfully",
	})
}

func GetAllProducts(e echo.Context) error {
	products, err := usecases.GetAllProducts()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to retrieve products", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"products": products,
	})
}

func GetProductByID(e echo.Context) error {
	idParam := e.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid product ID", "details": err.Error(),
		})
	}

	product, err := usecases.GetProductByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{
			"error": "Product not found", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, product)
}

func UpdateProduct(e echo.Context) error {
	idParam := e.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid product ID", "details": err.Error(),
		})
	}

	var updateData models.Product
	if err := e.Bind(&updateData); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body", "details": err.Error(),
		})
	}

	if updateData.Name == "" || updateData.Price <= 0 || updateData.Stock < 0 {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Name, price, and stock are required and must be valid",
		})
	}

	err = usecases.UpdateProduct(id, updateData)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to update product", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message": "Product updated successfully",
	})
}

func DeleteProduct(e echo.Context) error {
	idParam := e.Param("id")

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid product ID", "details": err.Error(),
		})
	}

	err = usecases.DeleteProduct(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to delete product", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message": "Product deleted successfully",
	})
}
