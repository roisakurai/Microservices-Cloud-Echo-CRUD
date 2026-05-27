package repositories

import (
	"context"
	"order-service/config"
	"order-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// helper untuk product dan user

func FindProductByID(ctx context.Context, id primitive.ObjectID) (models.Product, error) {
	var product models.Product
	err := config.DB.Collection("products").FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	return product, err
}

func UpdateProductStock(ctx context.Context, productID primitive.ObjectID, incrementValue int) error {
	_, err := config.DB.Collection("products").UpdateOne(ctx,
		bson.M{"_id": productID},
		bson.M{"$inc": bson.M{"stock": incrementValue}},
	)
	return err
}

func FindUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := config.DB.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

// crud order

func InsertOrder(ctx context.Context, order models.Order) error {
	_, err := config.DB.Collection("orders").InsertOne(ctx, order)
	return err
}

func FindAllOrders(ctx context.Context) ([]models.Order, error) {
	cursor, err := config.DB.Collection("orders").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func FindOrderByID(ctx context.Context, id primitive.ObjectID) (models.Order, error) {
	var order models.Order
	err := config.DB.Collection("orders").FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	return order, err
}

func UpdateOrder(ctx context.Context, id primitive.ObjectID, updateData bson.M) error {
	update := bson.M{"$set": updateData}
	result, err := config.DB.Collection("orders").UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func DeleteOrder(ctx context.Context, id primitive.ObjectID) error {
	result, err := config.DB.Collection("orders").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func UpdateOrderStatusCron() (int64, int64, error) {
	orderCollection := config.DB.Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	yesterday := time.Now().Add(-24 * time.Hour)

	filter := bson.M{
		"status": "pending",
		"created_at": bson.M{
			"$lte": yesterday,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"status":     "completed",
			"updated_at": time.Now(),
		},
	}

	result, err := orderCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, 0, err
	}

	return result.MatchedCount, result.ModifiedCount, nil
}

func CreateOrderIndexes() error {
	orderCollection := config.DB.Collection("orders")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "status", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "created_at", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "status", Value: 1},
				{Key: "created_at", Value: 1},
			},
		},
	}

	_, err := orderCollection.Indexes().CreateMany(ctx, indexes)
	return err
}
