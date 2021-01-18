/*
 * @File: database.mongo.go
 * @Description: Handles MongoDB connections
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package database defines connections with database
package database

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/Akshit8/go-boilerplate/config"
	"github.com/Akshit8/go-boilerplate/models"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	MgDBSession  *mgo.Session
	DatabaseName string
}

// Init initializes mongo database
func (db *MongoDB) Init() error {
	db.DatabaseName = config.Config.MgDBName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs: []string{
			"cluster0-shard-00-01.y2ty6.mongodb.net:27017",
			"cluster0-shard-00-00.y2ty6.mongodb.net:27017",
			"cluster0-shard-00-02.y2ty6.mongodb.net:27017",
		}, // Get HOST + PORT
		Timeout: 60 * time.Second,
		// Database: db.DatabaseName,            // Database name
		Username: config.Config.MgDBUsername, // Username
		Password: config.Config.MgDBPassword, // Password
	}

	tlsConfig := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig) // add TLS config
		return conn, err
	}

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.MgDBSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		log.Error("Can't connect to mongo", err)
	}

	return db.initData()
}

// InitData initializes default data
func (db *MongoDB) initData() error {
	var err error
	var count int

	// Check if user collection has at least one document
	sessionCopy := db.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(db.DatabaseName).C(config.ColUsers)
	count, err = collection.Find(bson.M{}).Count()

	if count < 1 {
		// Create admin/admin account
		user := models.User{ID: bson.NewObjectId(), Name: "admin", Password: "adminPassword"}
		err = collection.Insert(&user)
	}

	return err
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.MgDBSession != nil {
		log.Info("closing connection to mongo db")
		db.MgDBSession.Close()
	}
}
