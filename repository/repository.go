package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/Adi-Gupta018/react-mongo-crud-golang/model"
)

var (
	ErrCitizenNotFound = errors.New("citizen not found")
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) GetCitizen(ctx context.Context, id string) (model.Citizen, error) {
	var out model.Citizen
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Citizen{}, err
	}
	err = r.db.
		Collection("citizens").
		FindOne(ctx, bson.M{"_id": objectID}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Citizen{}, ErrCitizenNotFound
		}
		return model.Citizen{}, err
	}
	return out, nil
}

func (r *repository) GetAllCitizens(ctx context.Context) ([]model.Citizen, error) {
	cursor, err := r.db.
		Collection("citizens").
		Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var citizens []model.Citizen
	for cursor.Next(ctx) {
		var citizen model.Citizen
		if err := cursor.Decode(&citizen); err != nil {
			return nil, err
		}
		citizens = append(citizens, citizen)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return citizens, nil
}

func (r *repository) CreateCitizen(ctx context.Context, citizen model.Citizen) (model.Citizen, error) {
	citizen.ID = primitive.NewObjectID().Hex()
	_, err := r.db.
		Collection("citizens").
		InsertOne(ctx, citizen)
	if err != nil {
		return model.Citizen{}, err
	}
	return citizen, nil
}

func (r *repository) UpdateCitizen(ctx context.Context, citizen model.Citizen) (model.Citizen, error) {
	objectID, err := primitive.ObjectIDFromHex(citizen.ID)
	if err != nil {
		return model.Citizen{}, err
	}
	_, err = r.db.
		Collection("citizens").
		UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": citizen})
	if err != nil {
		return model.Citizen{}, err
	}
	return citizen, nil
}

func (r *repository) DeleteCitizen(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.db.
		Collection("citizens").
		DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrCitizenNotFound
	}
	return nil
}
