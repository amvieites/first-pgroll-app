package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEquipment creates a new Equipment record
func CreateEquipment(c *gin.Context) {
	var equipment Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&equipment)
	c.JSON(http.StatusCreated, equipment)
}

// GetEquipment retrieves an Equipment record by ID
func GetEquipment(c *gin.Context) {
	var equipment Equipment
	id := c.Param("id")

	if err := DB.First(&equipment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}
	c.JSON(http.StatusOK, equipment)
}

// GetEquipments retrieves all Equipment records
func GetEquipments(c *gin.Context) {
	var equipments []Equipment
	DB.Find(&equipments)
	c.JSON(http.StatusOK, equipments)
}

// UpdateEquipment updates an Equipment record by ID
func UpdateEquipment(c *gin.Context) {
	var equipment Equipment
	id := c.Param("id")

	if err := DB.First(&equipment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}

	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&equipment)
	c.JSON(http.StatusOK, equipment)
}

// DeleteEquipment deletes an Equipment record by ID
func DeleteEquipment(c *gin.Context) {
	var equipment Equipment
	id := c.Param("id")

	if err := DB.First(&equipment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}

	DB.Delete(&equipment)
	c.JSON(http.StatusNoContent, gin.H{})
}
