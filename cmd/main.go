package main

import (
	"context"
	"github.com/TeamStrata/strata/pkg/api"
	"github.com/TeamStrata/strata/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-set"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const initialUserLimit = 20

func main() {
	server := gin.Default()

	// Cors setup
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{
		"*",
	}
	server.Use(cors.New(config))

	// Get database connection string
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}
	conStr := os.Getenv("CONNECTION_STRING")

	// Init database manager
	db, err := database.NewDbManager(conStr, context.Background())
	if err != nil {
		log.Fatalf("error initializing DB manager: %s", err.Error())
	}

	users := set.New[string](initialUserLimit)

	// Endpoints
	server.POST("/login", api.LoginHandler(db, users))
	server.POST("/signup", api.SignUpHandler(db, users))
	server.GET("/ping", api.PingHandler)

	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
