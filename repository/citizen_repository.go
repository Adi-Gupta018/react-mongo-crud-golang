package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCitizens retrieves all citizens from the database
func (cr *CitizenRepository) GetCitizens(ctx context.Context) ([]Citizen, error) {
	var citizens []Citizen
	cursor, err := cr.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var citizen Citizen
		if err := cursor.Decode(&citizen); err != nil {
			return nil, err
		}
		citizens = append(citizens, citizen)
	}
	return citizens, nil
}

// GetCitizen retrieves a citizen by their ID from the database
func (cr *CitizenRepository) GetCitizen(ctx context.Context, id string) (Citizen, error) {
	var citizen Citizen
	err := cr.collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&citizen)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Citizen{}, ErrCitizenNotFound
		}
		return Citizen{}, err
	}
	return citizen, nil
}

// CreateCitizen creates a new citizen in the database
func (cr *CitizenRepository) CreateCitizen(ctx context.Context, citizen Citizen) (Citizen, error) {
	_, err := cr.collection.InsertOne(ctx, citizen)
	if err != nil {
		return Citizen{}, err
	}
	return citizen, nil
}

// UpdateCitizen updates an existing citizen in the database
func (cr *CitizenRepository) UpdateCitizen(ctx context.Context, citizen Citizen) (Citizen, error) {
	_, err := cr.collection.ReplaceOne(ctx, bson.D{{Key: "_id", Value: citizen.ID}}, citizen)
	if err != nil {
		return Citizen{}, err
	}
	return citizen, nil
}

// DeleteCitizen deletes a citizen by their ID from the database
func (cr *CitizenRepository) DeleteCitizen(ctx context.Context, id string) error {
	_, err := cr.collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
}
