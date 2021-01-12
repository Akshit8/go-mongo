/*
 * @File: models.error.go
 * @Description: Defines Error information that will be returned to the clients
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package models defines information that will be returned to the clients
package models

// Error defines the response error
type Error struct {
	Code    int    `json:"code" example:"27"`
	Message string `json:"message" example:"Error message"`
}