/*
 * @File: models.user.go
 * @Description: Defines User model
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package models defines information that will be returned to the clients
package models

import "gopkg.in/mgo.v2/bson"

// User information
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string        `bson:"name" json:"name" example:"raycad"`
	Email    string        `bson:"email" json:"email" example:"raycad@mail.com"`
	Password string        `bson:"password" json:"password" example:"raycad"`
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Email    string `json:"email" example:"raycad@mail.com"`
	Password string `json:"password" example:"User Password"`
}
