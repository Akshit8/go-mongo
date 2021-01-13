/*
 * @File: utils.utils.go
 * @Description: Reusable stuffs for services
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package utils for utilites
package utils

import (
	"errors"
	"time"

	"github.com/Akshit8/go-boilerplate/config"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

// Utils struct for binding GenerateJWT ValidateObjectID methods
type Utils struct{}

// GenerateJWT generates token from the given information
func (u *Utils) GenerateJWT(name string, role string) (string, error) {
	claims := SdtClaims{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    config.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.JwtSecretPassword))

	return tokenString, err
}

// ValidateObjectID checks the given ID if it's an object id or not
func (u *Utils) ValidateObjectID(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New(config.ErrNotObjectIDHex)
	}
	return nil
}
