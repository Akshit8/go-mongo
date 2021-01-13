/*
 * @File: models.user.go
 * @Description: Defines User model
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package models defines information that will be returned to the clients
package models

import (
	"errors"

	"github.com/Akshit8/go-boilerplate/config"
	"gopkg.in/mgo.v2/bson"
)

// User information
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id" example:"5bbdadf782ebac06a695a8e7"`
	Name     string        `bson:"name" json:"name" example:"raycad"`
	Password string        `bson:"password" json:"password" example:"raycad"`
}

// AddUser information
type AddUser struct {
	Name     string `json:"name" example:"User Name"`
	Password string `json:"password" example:"User Password"`
}

// Validate user
func (a AddUser) Validate() error {
	switch {
	case len(a.Name) == 0:
		return errors.New(config.ErrNameEmpty)
	case len(a.Password) == 0:
		return errors.New(config.ErrPasswordEmpty)
	default:
		return nil
	}
}
