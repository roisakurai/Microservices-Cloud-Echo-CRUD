package usecases

import (
	"context"
	"errors"
	"time"
	"user-service/models"
	"user-service/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// cek username
	usernameCount, err := repositories.CountByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if usernameCount > 0 {
		return errors.New("username already exists")
	}

	// cek email
	emailCount, err := repositories.CountByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if emailCount > 0 {
		return errors.New("email already exists")
	}

	// set timestamp waktu sekarang
	user.CreatedAt = time.Now()

	// memanggil repository
	return repositories.InsertUser(ctx, user)
}

func GetAllUsers() ([]models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := repositories.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var responses []models.UserResponse
	for _, user := range users {

		// Pengecekan Zero Value untuk UpdatedAt
		updatedAtStr := ""
		if !user.UpdatedAt.IsZero() {
			updatedAtStr = user.UpdatedAt.Format("02 January 2006 15:04:05")
		}

		responses = append(responses, models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("02 January 2006 15:04:05"),
			UpdatedAt: updatedAtStr, // akan berisi string kosong jika belum pernah diupdate
		})
	}

	return responses, nil
}

func GetUserByID(id primitive.ObjectID) (models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := repositories.FindUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	// pengecekan UpdatedAt
	updatedAtStr := ""
	if !user.UpdatedAt.IsZero() {
		updatedAtStr = user.UpdatedAt.Format("02 January 2006 15:04:05")
	}

	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("02 January 2006 15:04:05"),
		UpdatedAt: updatedAtStr,
	}

	return response, nil
}

func UpdateUser(id primitive.ObjectID, updateData models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usernameCount, err := repositories.CountByUsername(ctx, updateData.Username)
	if err != nil {
		return err
	}
	if usernameCount > 0 {
		return errors.New("username already exists")
	}

	emailCount, err := repositories.CountByEmail(ctx, updateData.Email)
	if err != nil {
		return err
	}
	if emailCount > 0 {
		return errors.New("email already exists")
	}

	// set data apa saja yang boleh diubah
	updateBson := bson.M{
		"username":   updateData.Username,
		"email":      updateData.Email,
		"updated_at": time.Now(),
	}

	return repositories.UpdateUser(ctx, id, updateBson)
}

func DeleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repositories.DeleteUser(ctx, id)
}
