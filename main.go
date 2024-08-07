package main

import (
	"log"
	"net/http"
	"os"

	"github.com/balasl342/kibana-go-example-logrus/logger"
	"github.com/balasl342/kibana-go-example-logrus/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve Elasticsearch config from environment variables
	esURL := os.Getenv("ELASTICSEARCH_URL")
	esAPIKey := os.Getenv("ELASTICSEARCH_API_KEY")

	// Initialize Elasticsearch hook
	hook, err := logger.NewElasticsearchHook(esURL, esAPIKey)
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch hook: %v", err)
	}

	// Initialize logger
	logger := logrus.New()
	logger.Hooks.Add(hook)

	// Set up router and routes
	router := mux.NewRouter()
	routes.SetupRoutes(router, logger)

	// Start the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
