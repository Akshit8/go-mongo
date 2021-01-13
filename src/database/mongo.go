/*
 * @File: database.mongo.go
 * @Description: Handles MongoDB connections
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package database defines connections with database
package database

import (
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
func (db *MongoDB) Init(addressString string) error {
	db.DatabaseName = config.Config.MgDBName

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.MgDBSession, err = mgo.Dial(addressString)

	if err != nil {
		log.Error("Can't connect to mongo, notes error: ", err)
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
		user := models.User{ID: bson.NewObjectId(), Name: "admin", Email: "admin@mail.com", Password: "adminPassword"}
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
