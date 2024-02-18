package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config corresponds to the fields in config.yaml
type Config struct {
	APIKey *string `yaml:"api_key"`
}

// LoadConfig loads `config.yaml` and unmarshals it into a struct
func LoadConfig() Config {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	c := Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
