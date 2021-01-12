/*
 * @File: database.mongo.go
 * @Description: Handles MongoDB connections
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package databases defines connections with database
package databases

import (
	"time"

	"github.com/Akshit8/go-boilerplate/config"
	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	MgDBSession  *mgo.Session
	DatabaseName string
}

func (db *MongoDB) Init() error {
	db.DatabaseName = config.Config.MgDBName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{config.Config.MgAddrs}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: db.DatabaseName,            // Database name
		Username: config.Config.MgDBUsername, // Username
		Password: config.Config.MgDBPassword, // Password
	}
}
