package data

import (
	"fmt"
	"net/http"
	"os"
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

	if !isAllowedExt(ext) {
		c.JSON(400, gin.H{"error": "bad file extension"})
		return
	}

	if !isAllowedFileSize(file.Size) {
		c.JSON(400, gin.H{"error": "bad file size"})
		return
	}

	if err = c.SaveUploadedFile(file, "uploads/"+data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", data),
	})
}

func GetData(c *gin.Context) {
	fileName := c.Param("file")

	file, err := os.Open("uploads/" + fileName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file info"})
		return
	}
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	c.Writer.Header().Set("Content-Type", http.DetectContentType([]byte(fileName)))
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	http.ServeContent(c.Writer, c.Request, fileName, fileInfo.ModTime(), file)
}

func DeleteData(c *gin.Context) {
	fileName := c.Param("file")

	err := os.Remove("uploads/" + fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}
