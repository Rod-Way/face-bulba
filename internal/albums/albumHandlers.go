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
