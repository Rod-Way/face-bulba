package data

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", file.Filename),
	})
}
