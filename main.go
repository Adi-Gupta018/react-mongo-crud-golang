package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env file error")
	}
	var conn_url = os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(conn_url)
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
