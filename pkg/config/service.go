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
}

func LoadEnv(rootPath string, config *Repository) {
	if err := godotenv.Load(filepath.Join(strings.TrimSpace(rootPath), ".env")); err != nil {
		log.Fatal(err)
	}
	config.VirtuosoServerURL = mustGetEnv("VIRTUOSO_SERVER_URL")
}

func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("environment variable %s not set", k)
	}
	return v
}
