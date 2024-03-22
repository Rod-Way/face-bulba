package albums

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAlbum(c *gin.Context) {
	var album = NewAlbum()
	if err := c.ShouldBindJSON(&album); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	album.Users = append(album.Users, c.GetString("username"))
	album.ID = primitive.NewObjectID()
	album.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	if err := album.SaveAlbum(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := AddAlbumToUser(album.ID, c.GetString("username")); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Album created successfully",
	})
}

func UpdateAlbum(c *gin.Context) {
	albumID := c.Param("id")

	var updateFields map[string]interface{}
	if err := c.BindJSON(&updateFields); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ID, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid albumID"})
		return
	}

	// Check author and user
	if albumUS, err := FindUser(ID, c.GetString("username")); err != nil {
		fmt.Println(c.GetString("username"), albumUS)
		c.JSON(500, gin.H{
			"error": "can not get album username",
		})
		return
	}

	for field, value := range updateFields {
		if !isAllowedField(field) {
			c.JSON(400, gin.H{"error": "field not allowed"})
			return
		}
		if err := UpdateField(ID, field, value); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"message": "Album updated successfully"})
}

func DeleteAlbum(c *gin.Context) {
	var albumID string
	if err := c.ShouldBindJSON(&albumID); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "invalid albumID",
		})
	}

	// Check author and user
	if albumUS, err := FindUser(objID, c.GetString("username")); err != nil {
		fmt.Println(c.GetString("username"), albumUS)
		c.JSON(500, gin.H{
			"error": "can not get album username",
		})
		return
	}

	if err := DeleteAlbumByID(objID); err != nil {
		c.JSON(500, gin.H{
			"error": "Can not delete",
		})
	}

	if err := DeleteAlbumFromUser(objID, c.GetString("username")); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "album deleted successfully",
	})
}

func GetAlbumByID(c *gin.Context) {
	albumID := c.Param("id")
	ID, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalig postID"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "invalid postID",
		})
	}

	// Check author and user
	author, err := FindUser(objID, "username")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "can not get album author username",
		})
		return
	}

	if author != c.GetString("username") {
		c.JSON(400, gin.H{
			"error": "invalid author",
		})
		return
	}

	album, err := GetAlbumByIDDB(ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": album,
	})
}
