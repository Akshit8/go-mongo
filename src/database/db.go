/*
 * @File: database.db.go
 * @Description: Creates global database instance
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

// Package database created global db interface
package database

// Database shares global database instance
var (
	Database MongoDB
)
