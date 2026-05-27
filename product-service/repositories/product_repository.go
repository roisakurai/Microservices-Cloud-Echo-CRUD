package repositories

import (
	"context"
	"product-service/config"
	"product-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertProduct(ctx context.Context, product models.Product) error {
	collection := config.DB.Collection("products")
	_, err := collection.InsertOne(ctx, product)
	return err
}

func FindAllProducts(ctx context.Context) ([]models.Product, error) {
	collection := config.DB.Collection("products")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func FindProductByID(ctx context.Context, id primitive.ObjectID) (models.Product, error) {
	collection := config.DB.Collection("products")

	var product models.Product
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)

	return product, err
}

func UpdateProduct(ctx context.Context, id primitive.ObjectID, updateData bson.M) error {
	collection := config.DB.Collection("products")

	update := bson.M{"$set": updateData}
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)

	return err
}

func DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	collection := config.DB.Collection("products")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
