package repository

import (
	"context"

	"github.com/Adi-Gupta018/react-mongo-crud-golang/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetCitizen(ctx context.Context, id primitive.ObjectID) (model.Citizen, error)
	GetAllCitizens(ctx context.Context) ([]model.Citizen, error) // New function added
	CreateCitizen(ctx context.Context, citizen model.Citizen) (model.Citizen, error)
	UpdateCitizen(ctx context.Context, citizen model.Citizen) (model.Citizen, error)
	DeleteCitizen(ctx context.Context, id primitive.ObjectID) error
}
