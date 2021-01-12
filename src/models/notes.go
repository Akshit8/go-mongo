/*
 * @File: models.note.go
 * @Description: Defines Note model
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package models defines information that will be returned to the clients
package models

import "gopkg.in/mgo.v2/bson"

// Note information
type Note struct {
	ID          bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Title       string        `bson:"title" json:"title" example:"sample title"`
	Description string        `bson:"description" json:"description" example:"sample description"`
	Completed   bool          `bson:"completed" json:"completed" example:"true"`
	Owner       bson.ObjectId `bson:"owner" json:"owner" example:"5bbdadf782ebac06a695a8e7"`
}

// AddNote information
type AddNote struct {
	Title       string        `bson:"title" json:"title" example:"sample title"`
	Description string        `bson:"description" json:"description" example:"sample description"`
	Completed   bool          `bson:"completed" json:"completed" example:"true"`
	Owner       bson.ObjectId `bson:"owner" json:"owner" example:"5bbdadf782ebac06a695a8e7"`
}
