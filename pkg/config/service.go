package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// GraphURI is used for type safety
type GraphURI string

// VirtuosoURL is used for type safety
type VirtuosoURL string

type Repository struct {
	VirtuosoURL      VirtuosoURL
	GraphURI         GraphURI
	TestGraphURI     GraphURI
	VirtuosoUsername string
	VirtuosoPassword string
}

func Load(rootPath string, config *Repository) {
	if err := godotenv.Load(filepath.Join(strings.TrimSpace(rootPath), ".env")); err != nil {
		log.Println("ignoring error when loading env file:", err)
	}

	config.VirtuosoURL = VirtuosoURL(mustGetENV("VIRTUOSO_SERVER_URL"))
	config.GraphURI = GraphURI(mustGetENV("VIRTUOSO_GRAPH_URI"))
	config.TestGraphURI = GraphURI(mustGetENV("VIRTUOSO_TEST_GRAPH_URI"))
	config.VirtuosoUsername = mustGetENV("VIRTUOSO_USERNAME")
	config.VirtuosoPassword = mustGetENV("VIRTUOSO_PASSWORD")
}

func mustGetENV(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing ENV key: %s\n", key)
	}
	return value
}
