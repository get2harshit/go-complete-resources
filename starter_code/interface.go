package database

import "fmt"

// Interface (contract)
type Database interface {
	Save()
}

// Mongo implementation
type Mongo struct{
	Uri string
}

// Task 1: create save method for Mongo

// MySQL implementation
type MySQL struct{
	Uri string
}

// Task 2: create save method for MySQL

// Business logic depends on interface, not concrete DB
func Process(db Database) {
  // Task 3: write business logic
}

Complete all 3 tasks
