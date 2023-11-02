package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type Repository struct {
	VirtuosoServerURL string `json:"virtuoso_server_url"`
	GraphURI          string `json:"graph_uri"`
	TestGraphURI      string `json:"test_graph_uri"`
}

func Load(rootPath string, config *Repository) {
	file, err := os.Open(filepath.Join(strings.TrimSpace(rootPath), "config.json"))
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}

	defer file.Close()

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("error parsing config file: %v", err)
	}

	// check if config type has all fields set that are marked with json tag
	configType := reflect.TypeOf(config).Elem()
	value := reflect.ValueOf(*config)
	for i := 0; i < configType.NumField(); i++ {
		fieldName := configType.Field(i).Name
		field := value.FieldByName(fieldName)
		if tag := configType.Field(i).Tag.Get("json"); tag != "" && field.String() == "" {
			log.Fatalf("Missing config key: %s (json: %s)\n", fieldName, tag)
		}
	}
}
