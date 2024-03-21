package data

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ext := filepath.Ext(file.Filename)

	data := primitive.NewObjectID().Hex()

	if err = c.SaveUploadedFile(file, "uploads/"+ext+data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", data),
	})
}
