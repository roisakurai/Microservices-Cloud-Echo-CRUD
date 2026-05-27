package usecases

import (
	"context"
	"errors"
	"order-service/models"
	"order-service/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateOrder(req models.OrderRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if req.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	product, err := repositories.FindProductByID(ctx, req.ProductID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("product not found")
		}
		return err
	}

	if product.Stock < req.Quantity {
		return errors.New("product stock is not enough")
	}

	// mengurangi stok produk sesuai dengan jumlah pesanan
	err = repositories.UpdateProductStock(ctx, req.ProductID, -req.Quantity)
	if err != nil {
		return err
	}

	order := models.Order{
		ID:        primitive.NewObjectID(),
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Total:     product.Price * float64(req.Quantity),
		Status:    "Pending",
		CreatedAt: time.Now(),
	}

	return repositories.InsertOrder(ctx, order)
}

func GetAllOrders() ([]models.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	orders, err := repositories.FindAllOrders(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.OrderResponse

	// mapping relasi user & product secara manual
	for _, order := range orders {
		user, err := repositories.FindUserByID(ctx, order.UserID)
		if err != nil {
			user.Username = "User not found"
		}

		product, err := repositories.FindProductByID(ctx, order.ProductID)
		if err != nil {
			product.Name = "Product not found"
			product.Price = 0
		}

		// mengecek zero value untuk UpdatedAt
		updatedAtStr := ""
		if !order.UpdatedAt.IsZero() {
			updatedAtStr = order.UpdatedAt.Format("02 January 2006 15:04:05")
		}

		response = append(response, models.OrderResponse{
			ID:           order.ID,
			UserID:       order.UserID,
			Username:     user.Username,
			ProductID:    order.ProductID,
			ProductName:  product.Name,
			ProductPrice: product.Price,
			Quantity:     order.Quantity,
			Total:        order.Total,
			Status:       order.Status,
			CreatedAt:    order.CreatedAt.Format("02 January 2006 15:04:05"),
			UpdatedAt:    updatedAtStr,
		})

	}

	return response, nil
}

func GetOrderByID(id primitive.ObjectID) (models.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := repositories.FindOrderByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.OrderResponse{}, errors.New("order not found")
		}
		return models.OrderResponse{}, err
	}

	user, err := repositories.FindUserByID(ctx, order.UserID)
	if err != nil {
		user.Username = "User not found"
	}

	product, err := repositories.FindProductByID(ctx, order.ProductID)
	if err != nil {
		product.Name = "Product not found"
		product.Price = 0
	}

	updatedAtStr := ""
	if !order.UpdatedAt.IsZero() {
		updatedAtStr = order.UpdatedAt.Format("02 January 2006 15:04:05")
	}

	response := models.OrderResponse{
		ID:           order.ID,
		UserID:       order.UserID,
		Username:     user.Username,
		ProductID:    order.ProductID,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     order.Quantity,
		Total:        order.Total,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt.Format("02 January 2006 15:04:05"),
		UpdatedAt:    updatedAtStr,
	}

	return response, nil
}

func UpdateOrder(id primitive.ObjectID, updateData models.OrderRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if updateData.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	oldOrder, err := repositories.FindOrderByID(ctx, id)
	if err != nil {
		return errors.New("order not found")
	}

	newProduct, err := repositories.FindProductByID(ctx, updateData.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	// produknya masih sama, hanya ganti jumlah (quantity)
	if oldOrder.ProductID == updateData.ProductID {
		difference := updateData.Quantity - oldOrder.Quantity

		if difference > 0 { // user menambah pesanan
			if newProduct.Stock < difference {
				return errors.New("product stock is not enough")
			}
			err = repositories.UpdateProductStock(ctx, updateData.ProductID, -difference)
		} else if difference < 0 { // user mengurangi pesanan
			stockBack := -difference // mengubah jadi positif
			err = repositories.UpdateProductStock(ctx, updateData.ProductID, stockBack)
		}

		if err != nil {
			return err
		}
	} else {
		// produknya diganti produk lain
		// mengembalikan stok produk lama
		err = repositories.UpdateProductStock(ctx, oldOrder.ProductID, oldOrder.Quantity)
		if err != nil {
			return err
		}

		// mengecek stok produk baru
		if newProduct.Stock < updateData.Quantity {
			return errors.New("new product stock is not enough")
		}

		// mengurangi stok produk baru
		err = repositories.UpdateProductStock(ctx, updateData.ProductID, -updateData.Quantity)
		if err != nil {
			return err
		}
	}

	// perhitungan total harga berdasarkan produk baru dan quantity baru
	total := newProduct.Price * float64(updateData.Quantity)

	updateBson := bson.M{
		"user_id":    updateData.UserID,
		"product_id": updateData.ProductID,
		"quantity":   updateData.Quantity,
		"total":      total,
		"updated_at": time.Now(),
	}

	err = repositories.UpdateOrder(ctx, id, updateBson)
	if err == mongo.ErrNoDocuments {
		return errors.New("order not found")
	}
	return err
}

func DeleteOrder(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := repositories.FindOrderByID(ctx, id)
	if err != nil {
		return errors.New("order not found")
	}

	// Kembalikan stok produk ke semula
	err = repositories.UpdateProductStock(ctx, order.ProductID, order.Quantity)
	if err != nil {
		return err
	}

	err = repositories.DeleteOrder(ctx, id)
	if err == mongo.ErrNoDocuments {
		return errors.New("order not found")
	}
	return err
}

func UpdateOrderStatusCron() (int64, int64, error) {
	return repositories.UpdateOrderStatusCron()
}
