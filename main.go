package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Adi-Gupta018/react-mongo-crud-golang/http"
	"github.com/Adi-Gupta018/react-mongo-crud-golang/repository"
)

func main() {
	// Create a MongoDB client
	client := db()
	defer client.Disconnect(context.TODO())

	// Create a repository
	repo := repository.NewRepository(client.Database("citizens"))

	// Create an HTTP server
	server := http.NewServer(repo)

	// Create a Gin router
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	{
		router.GET("/citizens/:id", server.GetCitizen)
		router.GET("/citizens/:id/", server.GetCitizen)
		router.GET("/citizens", server.GetAllCitizens) // New route added for getting all citizens
		router.GET("/citizens/", server.GetAllCitizens)
		router.POST("/citizens", server.CreateCitizen)
		router.POST("/citizens/", server.CreateCitizen)
		router.PUT("/citizens/:id", server.UpdateCitizen)
		router.PUT("/citizens/:id/", server.UpdateCitizen)
		router.DELETE("/citizens/:id", server.DeleteCitizen)
		router.DELETE("/citizens/:id/", server.DeleteCitizen)
	}

	// Start the router
	router.Run(":9080")
}

func db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Demo03:Demomongo03@demo.antwhfs.mongodb.net/?retryWrites=true&w=majority&appName=Demo")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}
