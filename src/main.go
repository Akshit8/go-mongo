/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package main

import (
	"io"
	"os"

	"github.com/Akshit8/go-boilerplate/config"
	"github.com/Akshit8/go-boilerplate/database"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	mongoDBAddressString := "mongodb+srv://akshit:akshit2853@cluster0.y2ty6.mongodb.net/<dbname>?retryWrites=true&w=majority"

	// Initialize User database
	err = database.Database.Init(mongoDBAddressString)
	if err != nil {
		return err
	}

	if config.Config.EnableGinConsoleLog {
		gin.DefaultWriter = io.MultiWriter(os.Stdout)
	} else {
		gin.DefaultWriter = io.MultiWriter()
	}

	m.router = gin.Default()

	return nil
}

// @title NotesManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /api/v1
func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	defer database.Database.Close()

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(config.Config.Port)
}
