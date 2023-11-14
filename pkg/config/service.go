package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type Repository struct {
	VirtuosoServerURL string
	GraphURI          string
	TestGraphURI      string
}

func Load(rootPath string, config *Repository) {
	if err := godotenv.Load(filepath.Join(strings.TrimSpace(rootPath), ".env")); err != nil {
		log.Println("ignoring error when loading env file:", err)
	}

	config.VirtuosoServerURL = mustGetENV("VIRTUOSO_SERVER_URL")
	config.GraphURI = mustGetENV("VIRTUOSO_GRAPH_URI")
	config.TestGraphURI = mustGetENV("VIRTUOSO_TEST_GRAPH_URI")
}

func mustGetENV(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing ENV key: %s\n", key)
	}
	return value
}
