/*
 * @File: models.error.go
 * @Description: Defines Message information that will be returned to the clients
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package models defines information that will be returned to the clients
package models

// Message defines the response message
type Message struct {
	Message string `json:"message" example:"message"`
}
