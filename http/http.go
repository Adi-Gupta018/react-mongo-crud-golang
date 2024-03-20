package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Adi-Gupta018/react-mongo-crud-golang/model"
	"github.com/Adi-Gupta018/react-mongo-crud-golang/repository"
)

type Server struct {
	repository repository.Repository
}

func NewServer(repository repository.Repository) *Server {
	return &Server{repository: repository}
}

func (s *Server) GetCitizen(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	ObjectID, err := primitive.ObjectIDFromHex(id)
	citizen, err := s.repository.GetCitizen(ctx, ObjectID)
	if err != nil {
		if errors.Is(err, repository.ErrCitizenNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"citizen": citizen})
}

func (s *Server) GetAllCitizens(ctx *gin.Context) {
	citizens, err := s.repository.GetAllCitizens(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"citizens": citizens})
}

func (s *Server) CreateCitizen(ctx *gin.Context) {
	var citizen model.Citizen
	if err := ctx.ShouldBindJSON(&citizen); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	citizen, err := s.repository.CreateCitizen(ctx, citizen)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"citizen": citizen})
}

func (s *Server) UpdateCitizen(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	ObjectID, _ := primitive.ObjectIDFromHex(id)
	var citizen model.Citizen
	if err := ctx.ShouldBindJSON(&citizen); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	citizen.ID = ObjectID
	citizen, err := s.repository.UpdateCitizen(ctx, citizen)
	if err != nil {
		if errors.Is(err, repository.ErrCitizenNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"citizen": citizen})
}

func (s *Server) DeleteCitizen(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	ObjectID, _ := primitive.ObjectIDFromHex(id)
	if err := s.repository.DeleteCitizen(ctx, ObjectID); err != nil {
		if errors.Is(err, repository.ErrCitizenNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
