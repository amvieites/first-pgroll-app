package main

import "github.com/gin-gonic/gin"

func main() {
	// Initialize database connection
	InitDB()

	// Set up Gin router
	r := gin.Default()

	// Equipment routes
	r.POST("/equipment", CreateEquipment)
	r.GET("/equipment/:id", GetEquipment)
	r.GET("/equipment", GetEquipments)
	r.PUT("/equipment/:id", UpdateEquipment)
	r.DELETE("/equipment/:id", DeleteEquipment)

	// Start the server
	r.Run(":8880")
}
