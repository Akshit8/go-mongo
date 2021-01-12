/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package main

import (
	"github.com/Akshit8/go-boilerplate/config"
	"github.com/gin-gonic/gin"
)

// Main manages main golang application
type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = config.LoadConfig()
	if err != nil {
		return err
	}
	return nil
}
