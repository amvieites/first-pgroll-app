package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Equipment struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

var DB *gorm.DB

func InitDB() {
	// Set up PostgreSQL connection
	dsn := "host=localhost user=t5r6qf password=xau_MkLaRvFFoOWmP6kTV50lwjD1jUbVtgMz1 dbname=first-pgroll-app port=7654 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the schema
	DB.AutoMigrate(&Equipment{})
}
