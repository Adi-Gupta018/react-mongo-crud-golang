package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// ErrCitizenNotFound represents the error when a citizen is not found in the database
var ErrCitizenNotFound = errors.New("citizen not found")

// Citizen represents the structure of a citizen entity
type Citizen struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty" bson:"dateOfBirth,omitempty"`
	Gender      string `json:"gender,omitempty" bson:"gender,omitempty"`
	Address     string `json:"address,omitempty" bson:"address,omitempty"`
	City        string `json:"city,omitempty" bson:"city,omitempty"`
	State       string `json:"state,omitempty" bson:"state,omitempty"`
	Pincode     string `json:"pincode,omitempty" bson:"pincode,omitempty"`
}

// Repository represents the interface for managing citizen data
type Repository interface {
	GetCitizens(ctx context.Context) ([]Citizen, error)
	GetCitizen(ctx context.Context, id string) (Citizen, error)
	CreateCitizen(ctx context.Context, citizen Citizen) (Citizen, error)
	UpdateCitizen(ctx context.Context, citizen Citizen) (Citizen, error)
	DeleteCitizen(ctx context.Context, id string) error
}

// CitizenRepository represents the repository for managing citizen data
type CitizenRepository struct {
	collection *mongo.Collection
}

// NewCitizenRepository creates a new instance of CitizenRepository
func NewCitizenRepository(collection *mongo.Collection) *CitizenRepository {
	return &CitizenRepository{collection: collection}
}
