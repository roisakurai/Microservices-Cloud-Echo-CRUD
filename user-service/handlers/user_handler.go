package handlers

import (
	"net/http"
	"user-service/models"
	"user-service/usecases"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(e echo.Context) error {
	var user models.User

	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body", "details": err.Error(),
		})
	}

	if user.Username == "" || user.Email == "" {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Username and email are required",
		})
	}

	// memanggil usecase untuk membuat user
	err := usecases.CreateUser(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create user", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, echo.Map{
		"message": "User created successfully",
	})
}

func GetAllUsers(e echo.Context) error {
	users, err := usecases.GetAllUsers()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to retrieve users", "details": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"users": users,
	})
}

func GetUserByID(e echo.Context) error {
	idParam := e.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID", "details": err.Error()})
	}

	user, err := usecases.GetUserByID(id)
	if err != nil {
		return e.JSON(http.StatusNotFound, echo.Map{"error": "User not found", "details": err.Error()})
	}

	return e.JSON(http.StatusOK, user)
}

func UpdateUser(e echo.Context) error {
	idParam := e.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID", "details": err.Error()})
	}

	var updateData models.User
	if err := e.Bind(&updateData); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body", "details": err.Error()})
	}

	if updateData.Username == "" || updateData.Email == "" {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"error": "Username and email are required",
		})
	}

	err = usecases.UpdateUser(id, updateData)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update user", "details": err.Error()})
	}

	return e.JSON(http.StatusOK, echo.Map{"message": "User updated successfully"})
}

func DeleteUser(e echo.Context) error {
	idParam := e.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user ID", "details": err.Error()})
	}

	err = usecases.DeleteUser(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete user", "details": err.Error()})
	}

	return e.JSON(http.StatusOK, echo.Map{"message": "User deleted successfully"})
}
