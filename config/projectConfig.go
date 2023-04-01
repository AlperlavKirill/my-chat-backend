package config

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port        int `yaml:"port" validate:"required,number"`
		ReservePort int `yaml:"reservePort" validate:"required,number"`
	} `yaml:"server"`
}

func Properties() *Config {
	yamlFile, err := os.Open("application.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
		return nil
	}
	defer yamlFile.Close()

	var cfg Config
	err = yaml.NewDecoder(yamlFile).Decode(&cfg)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
		return nil
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		log.Fatalf("Failed to validate config: %v", err)
		return nil
	}

	return &cfg
}
