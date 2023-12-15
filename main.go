// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string `json:"name"`
}

// Database connection
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("file:data.db?cache=shared&_loc=auto"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Auto Migrate the time_log table
	db.AutoMigrate(&Item{})
}

func main() {
	r := gin.Default()

	// api endpoints
	r.POST("/create", CreateHandler)
	r.GET("/read", ReadHandler)
	r.PUT("/update/:id", UpdateHandler)
	r.DELETE("/delete/:id", DeleteHandler)

	// Run the server
	r.Run(":7575")
}

func CreateHandler(c *gin.Context) {
	var newItem Item
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	// Create the new Item in the database
	db.Create(&newItem)
	c.JSON(http.StatusCreated, newItem)
}

func ReadHandler(c *gin.Context) {
	var Items []Item
	// Query all Item from the database
	db.Find(&Items)
	c.JSON(http.StatusOK, Items)
}

func UpdateHandler(c *gin.Context) {
	// Extract Item ID from the request parameters
	id := c.Param("id")
	// Parse the Item ID into an integer
	var ItemID int
	fmt.Sscan(id, &ItemID)
	var updatedItem Item
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	// Update the Item in the database
	db.Model(&Item{}).Where("id = ?", ItemID).Updates(updatedItem)
	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func DeleteHandler(c *gin.Context) {
	// Extract Item ID from the request parameters
	idParam := c.Param("id")
	// Parse the Item ID into an integer
	ItemID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Item ID"})
		return
	}
	// Check if the Item exists
	var existingItem Item
	result := db.First(&existingItem, ItemID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	db.Delete(&existingItem)
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
