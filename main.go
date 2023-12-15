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

type Student struct {
	gorm.Model
	Name    string `json:"name"`
	Course  string `json:"course"`
	College string `json:"college"`
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
	db.AutoMigrate(&Student{})
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
	var newStudent Student
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	// Create the new student in the database
	db.Create(&newStudent)
	c.JSON(http.StatusCreated, newStudent)
}

func ReadHandler(c *gin.Context) {
	var Students []Student
	// Query all student from the database
	db.Find(&Students)
	c.JSON(http.StatusOK, Students)
}

func UpdateHandler(c *gin.Context) {
	// Extract Student ID from the request parameters
	id := c.Param("id")
	// Parse the Student ID into an integer
	var StudentID int
	fmt.Sscan(id, &StudentID)
	var updatedStudent Student
	if err := c.BindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	// Update the student in the database
	db.Model(&Student{}).Where("id = ?", StudentID).Updates(updatedStudent)
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func DeleteHandler(c *gin.Context) {
	// Extract student ID from the request parameters
	idParam := c.Param("id")
	// Parse the student ID into an integer
	StudentID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}
	// Check if the student exists
	var existingStudent Student
	result := db.First(&existingStudent, StudentID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	db.Delete(&existingStudent)
	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}
