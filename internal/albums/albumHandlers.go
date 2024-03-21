package albums

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAlbum(c *gin.Context) {
	var comment = NewAlbum()
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	comment.Users = append(comment.Users, c.GetString("username"))
	comment.ID = primitive.NewObjectID()
	comment.CreatedAt = time.Now()

	if err := comment.SaveAlbum(); err != nil {
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
		c.JSON(400, gin.H{"error": "invalid postID"})
		return
	}

	// Check author and user
	author, err := GetField(ID, "username")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "can not get post username",
		})
		return
	}

	auth, ok := author.(string)
	if !ok || auth != c.GetString("username") {
		c.JSON(400, gin.H{
			"error": "invalid author",
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
	var postID string
	if err := c.ShouldBindJSON(&postID); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "invalid postID",
		})
	}

	// Check author and user
	author, err := GetField(objID, "username")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "can not get album author username",
		})
		return
	}

	auth, ok := author.(string)
	if !ok || auth != c.GetString("username") {
		c.JSON(400, gin.H{
			"error": "invalid author",
		})
		return
	}

	if err := DeleteAlbumByID(objID); err != nil {
		c.JSON(500, gin.H{
			"error": "Can not delete",
		})
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
