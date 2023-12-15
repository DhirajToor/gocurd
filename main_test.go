// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/create", CreateHandler)
	r.GET("/read", ReadHandler)
	r.PUT("/update/:id", UpdateHandler)
	r.DELETE("/delete/:id", DeleteHandler)

	return r
}

func TestCreateHandler(t *testing.T) {
	router := setupRouter()
	// Create a sample Item for testing
	newItem := Item{Name: "TestItem"}
	newItemJSON, err := json.Marshal(newItem)
	assert.NoError(t, err)
	req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(newItemJSON))
	assert.NoError(t, err)

	// Set the content type to JSON
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	var createdItem Item
	err = json.Unmarshal(w.Body.Bytes(), &createdItem)
	assert.NoError(t, err)
}

func TestReadHandler(t *testing.T) {
	router := setupRouter()

	// use /read endpoint
	req, err := http.NewRequest(http.MethodGet, "/read", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var Items []Item
	err = json.Unmarshal(w.Body.Bytes(), &Items)
	assert.NoError(t, err)
}

func TestUpdateHandler(t *testing.T) {
	router := setupRouter()

	newItem := Item{Name: "TestItem"}
	newItemJSON, err := json.Marshal(newItem)
	assert.NoError(t, err)

	// use /create endpoint to create a new Item for updating
	createReq, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(newItemJSON))
	assert.NoError(t, err)
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)
	assert.Equal(t, http.StatusCreated, createW.Code)

	var createdItem Item
	err = json.Unmarshal(createW.Body.Bytes(), &createdItem)
	assert.NoError(t, err)
	createdItem.Name = "TestItemUpdated"
	updatedItemJSON, err := json.Marshal(createdItem)
	assert.NoError(t, err)

	// Perform a PUT request to the /update/:id endpoint
	updateReq, err := http.NewRequest(http.MethodPut, "/update/"+strconv.Itoa(int(createdItem.ID)), bytes.NewBuffer(updatedItemJSON))
	assert.NoError(t, err)
	updateReq.Header.Set("Content-Type", "application/json")
	updateW := httptest.NewRecorder()
	router.ServeHTTP(updateW, updateReq)
	assert.Equal(t, http.StatusOK, updateW.Code)
}

func TestDeleteHandler(t *testing.T) {
	router := setupRouter()

	// Create a sample Item for testing
	newItem := Item{Name: "TestItemUpdate"}
	newItemJSON, err := json.Marshal(newItem)
	assert.NoError(t, err)

	// use /create to get stduent endpoint
	createReq, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(newItemJSON))
	assert.NoError(t, err)
	createReq.Header.Set("Content-Type", "application/json")

	createW := httptest.NewRecorder()
	router.ServeHTTP(createW, createReq)
	assert.Equal(t, http.StatusCreated, createW.Code)
	var createdItem Item
	err = json.Unmarshal(createW.Body.Bytes(), &createdItem)
	assert.NoError(t, err)

	// pass previous get id to /delete/:id endpoint
	deleteReq, err := http.NewRequest(http.MethodDelete, "/delete/"+strconv.Itoa(int(createdItem.ID)), nil)
	assert.NoError(t, err)
	deleteW := httptest.NewRecorder()
	router.ServeHTTP(deleteW, deleteReq)
	assert.Equal(t, http.StatusOK, deleteW.Code)
}
