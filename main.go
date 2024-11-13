// main.go
package main

import (
    "context"
    "log"
    "net/http"
    "os"

    "vlab-backend/handlers"
    "vlab-backend/middleware"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

// Middleware untuk mengatasi CORS
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

func main() {
    // Load environment variables
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize MongoDB client
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
    if err != nil {
        log.Fatal(err)
    }
    userCollection = client.Database("salsaDB").Collection("stateRecord")
    handlers.SetUserCollection(userCollection)

    // Setup Gin router
    router := gin.Default()

    // Use CORS middleware
    router.Use(CORSMiddleware())

    // Define routes
    router.POST("/register", handlers.Register)
    router.POST("/login", handlers.Login)
    router.GET("/userstate", middleware.Authenticate, handlers.UserState)
    router.POST("/userstate", middleware.Authenticate, handlers.UserState)

    // Start server
    router.Run(":8000")
}
