/*
 * @File: config.config.go
 * @Description: Defines configuration of the service
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package config defines service configuration
package config

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

// Configuration stores setting values
type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`

	MgAddrs      string `json:"mgAddrs"`
	MgDBName     string `json:"mgDbName"`
	MgDBUsername string `json:"mgDbUsername"`
	MgDBPassword string `json:"mgDbPassword"`

	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`
}

// Config shares the global configuration
var Config *Configuration

// COLLECTIONs of the database table
const (
	ColUsers = "users"
	ColNotes = "notes"
)

// Status messages
const (
	ErrNameEmpty      = "name is empty"
	ErrPasswordEmpty  = "password is empty"
	ErrNotObjectIDHex = "string is not a valid hex representation of an ObjectId"
)

// Status Codes
const (
	StatusCodeUnknown = -1
	StatusCodeOK      = 1000
	StatusMismatch    = 10
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Error(err)
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Service Logger
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the debug severity or above.
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}

