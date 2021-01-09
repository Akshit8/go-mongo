package db

import (
	"database/sql"

	"github.com/Akshit8/go-boilerplate/models"
)

// GetAllItems function
func (db Database) GetAllItems() (*models.ItemList, error) {
	list := &models.ItemList{}
	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Items = append(list.Items, item)
	}
	return list, nil
}

// AddItem function
func (db Database) AddItem(item *models.Item) error {
	var id int
	var createdAt string
	query := `INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, item.Name, item.Description).Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	item.ID = id
	item.CreatedAt = createdAt
	return nil
}

// GetItemByID function
func (db Database) GetItemByID(itemID int) (models.Item, error) {
	item := models.Item{}
	query := `SELECT * FROM items WHERE id = $1;`
	row := db.Conn.QueryRow(query, itemID)
	switch err := row.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt); err {
	case sql.ErrNoRows:
		return item, ErrNoMatch
	default:
		return item, err
	}
}

// DeleteItem function
func (db Database) DeleteItem(itemID int) error {
	query := `DELETE FROM items WHERE id = $1;`
	_, err := db.Conn.Exec(query, itemID)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

// UpdateItem function
func (db Database) UpdateItem(itemID int, itemData models.Item) (models.Item, error) {
	item := models.Item{}
	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
	err := db.Conn.QueryRow(query, itemData.Name, itemData.Description, itemID).Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}
	return item, nil
}
