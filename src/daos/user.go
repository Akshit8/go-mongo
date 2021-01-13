/*
 * @File: daos.user.go
 * @Description: Implements User CRUD functions for MongoDB
 * @Author: Akshit Sadana
 */

// Package daos for CRUD function with persistence layer
package daos

import (
	"github.com/Akshit8/go-boilerplate/config"
	"github.com/Akshit8/go-boilerplate/database"
	"github.com/Akshit8/go-boilerplate/models"
	"github.com/Akshit8/go-boilerplate/utils"
	"gopkg.in/mgo.v2/bson"
)

// User manages User CRUD
type User struct {
	utils *utils.Utils
}

// GetAll gets the list of Users
func (u *User) GetAll() ([]models.User, error) {
	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	var users []models.User
	err := collection.Find(bson.M{}).All(users)
	return users, err
}

// GetByID finds a User by its id
func (u *User) getByID(id string) (models.User, error) {
	var err error

	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return models.User{}, err
	}

	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	var user models.User
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// DeleteByID finds a User by its id
func (u *User) DeleteByID(id string) error {
	var err error

	err = u.utils.ValidateObjectID(id)
	if err != nil {
		return err
	}

	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// Login User
func (u *User) Login(name string, password string) (models.User, error) {
	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	var user models.User
	err := collection.Find(bson.M{
		"$and": []bson.M{
			{"name": name},
			{"password": password},
		}}).One(&user)
	return user, err
}

// Insert adds a new User into database'
func (u *User) Insert(user models.User) error {
	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	err := collection.Insert(&user)
	return err
}

// Delete remove an existing User
func (u *User) Delete(user models.User) error {
	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	err := collection.Remove(&user)
	return err
}

// Update modifies an existing User
func (u *User) Update(user models.User) error {
	sessionCopy := database.Database.MgDBSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(database.Database.DatabaseName).C(config.ColUsers)

	err := collection.UpdateId(user.ID, &user)
	return err
}
