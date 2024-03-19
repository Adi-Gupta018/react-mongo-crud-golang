package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUsers retrieves all users from the database
func (cr *CitizenRepository) GetUsers(ctx context.Context) ([]User, error) { // Change function signature
	var users []User // Change variable type
	cursor, err := cr.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User // Change variable type
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUser retrieves a user by their ID from the database
func (cr *CitizenRepository) GetUser(ctx context.Context, id string) (User, error) { // Change function signature
	var user User // Change variable type
	err := cr.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, ErrUserNotFound // Update error
		}
		return User{}, err
	}
	return user, nil
}

// CreateUser creates a new user in the database
func (cr *CitizenRepository) CreateUser(ctx context.Context, user User) (User, error) { // Change function signature and parameter type
	_, err := cr.collection.InsertOne(ctx, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func (cr *CitizenRepository) UpdateUser(ctx context.Context, user User) (User, error) { // Change function signature and parameter type
	_, err := cr.collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: user.ID}}, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// DeleteUser deletes a user by their ID from the database
func (cr *CitizenRepository) DeleteUser(ctx context.Context, id string) error {
	_, err := cr.collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
}
