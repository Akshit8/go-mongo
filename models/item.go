// Package models is ...
package models

import (
	"fmt"
	"net/http"
)

// Item struct
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// ItemList struct
type ItemList struct {
	Items []Item `json:"items"`
}

// Bind function
func (i *Item) Bind(r *http.Request) error {
	if i.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}

// Render function ItemList
func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Render function for Item
func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
