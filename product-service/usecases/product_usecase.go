package usecases

import (
	"context"
	"product-service/models"
	"product-service/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(product models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product.CreatedAt = time.Now()

	return repositories.InsertProduct(ctx, product)
}

func GetAllProducts() ([]models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	products, err := repositories.FindAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	var responses []models.ProductResponse
	for _, product := range products {
		updatedAtStr := ""
		if !product.UpdatedAt.IsZero() {
			updatedAtStr = product.UpdatedAt.Format("02 January 2006 15:04:05")
		}

		responses = append(responses, models.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			CreatedAt: product.CreatedAt.Format("02 January 2006 15:04:05"),
			UpdatedAt: updatedAtStr,
		})
	}

	return responses, nil
}

func GetProductByID(id primitive.ObjectID) (models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	product, err := repositories.FindProductByID(ctx, id)
	if err != nil {
		return models.ProductResponse{}, err
	}

	updatedAtStr := ""
	if !product.UpdatedAt.IsZero() {
		updatedAtStr = product.UpdatedAt.Format("02 January 2006 15:04:05")
	}

	response := models.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt.Format("02 January 2006 15:04:05"),
		UpdatedAt: updatedAtStr,
	}

	return response, nil
}

func UpdateProduct(id primitive.ObjectID, updateData models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perbaikan: Gunakan updated_at, bukan menimpa created_at
	updateBson := bson.M{
		"name":       updateData.Name,
		"price":      updateData.Price,
		"stock":      updateData.Stock,
		"updated_at": time.Now(),
	}

	return repositories.UpdateProduct(ctx, id, updateBson)
}

func DeleteProduct(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return repositories.DeleteProduct(ctx, id)
}
