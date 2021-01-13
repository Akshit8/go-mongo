/*
 * @File: main.go
 * @Description: Creates HTTP server & API groups of the UserManagement Service
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Akshit8/go-boilerplate/config"
	"github.com/Akshit8/go-boilerplate/controllers"
	"github.com/Akshit8/go-boilerplate/database"
	_ "github.com/Akshit8/go-boilerplate/docs"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	// Initialize User database
	err = database.Database.Init()
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

// @title User Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /api/v1
func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		fmt.Print("init server error")
		return
	}

	defer database.Database.Close()

	c := controllers.User{}
	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/users")

		// APIs need to be authenticated
		user.Use(jwt.Auth(config.Config.JwtSecretPassword))
		{
			user.POST("", c.AddUser)
			user.GET("/list", c.ListUsers)
			user.GET("detail/:id", c.GetUserByID)
			user.GET("/", c.GetUserByParams)
			user.DELETE(":id", c.DeleteUserByID)
			user.PATCH("", c.UpdateUser)
		}
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := m.router.Run(config.Config.Port)
	if err == nil {
		log.Info("server listening at port 8080")
	} else {
		log.Error("error starting server")
	}
}
