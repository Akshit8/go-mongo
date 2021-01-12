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
)

// Configuration stores setting values
type Configuration struct {
	Port string `json:"port"`

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

// Http Status messages
const (
	ErrNameEmpty      = "Name is empty"
	ErrPasswordEmpty  = "Password is empty"
	ErrNotObjectIDHex = "String is not a valid hex representation of an ObjectId"
)

// Http Status Codes
const (
	StatusCodeOk           = 200
	StatusCodeCreated      = 201
	StatusCodeBadRequest   = 400
	StatusCodeUnauthorized = 401
	StatusCodeForbidden    = 403
	StatusCodeNotFound     = 404
	StatusCodeServerError  = 500
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
