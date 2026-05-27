package repositories

import (
	"context"
	"user-service/config"
	"user-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CountByUsername(ctx context.Context, username string) (int64, error) {
	collection := config.DB.Collection("users")
	return collection.CountDocuments(ctx, bson.M{"username": username})
}

func CountByEmail(ctx context.Context, email string) (int64, error) {
	collection := config.DB.Collection("users")
	return collection.CountDocuments(ctx, bson.M{"email": email})
}

func InsertUser(ctx context.Context, user models.User) error {
	collection := config.DB.Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func FindAllUsers(ctx context.Context) ([]models.User, error) {
	collection := config.DB.Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func FindUserByID(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	collection := config.DB.Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func UpdateUser(ctx context.Context, id primitive.ObjectID, updateData bson.M) error {
	collection := config.DB.Collection("users")
	update := bson.M{"$set": updateData}
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

func DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	collection := config.DB.Collection("users")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
