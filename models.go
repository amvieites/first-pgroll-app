package main

import (
	"database/sql/driver"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Equipment struct {
	ID       uint          `gorm:"primaryKey" json:"id"`
	Type     EquipmentType `db:"type" json:"type"` // Use the custom EquipmentType
	Name     string        `json:"name"`
	ImageURL string        `json:"image_url"`
} // Define the custom type for the PostgreSQL ENUM
type EquipmentType string

// Define constants for each of the ENUM values
const (
	Carabiner EquipmentType = "carabiner"
	Quickdraw EquipmentType = "quickdraw"
	HMS       EquipmentType = "hms"
	Rope      EquipmentType = "rope"
	Other     EquipmentType = "other"
)

// List of all valid EquipmentType values
var AllEquipmentTypes = []EquipmentType{
	Carabiner, Quickdraw, HMS, Rope, Other,
}

// Implement the Scanner interface to read ENUM values from the database
func (e *EquipmentType) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan EquipmentType: %v", value)
	}

	// Validate that the value is a valid ENUM
	for _, enumValue := range AllEquipmentTypes {
		if strValue == string(enumValue) {
			*e = EquipmentType(strValue)
			return nil
		}
	}

	*e = Other
	return nil
}

// Implement the Valuer interface to save the ENUM value to the database
func (e EquipmentType) Value() (driver.Value, error) {
	// Ensure the value is valid before inserting it into the database
	for _, enumValue := range AllEquipmentTypes {
		if e == enumValue {
			return string(e), nil
		}
	}

	return string(Other), nil
}

var DB *gorm.DB

func InitDB() {
	// Set up PostgreSQL connection
	dsn := "host=localhost user=equipment password=equipment dbname=equipment port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the schema
	DB.AutoMigrate(&Equipment{})
}
