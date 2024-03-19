package posts

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context) {
	var post = NewPost()
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	post.ID = primitive.NewObjectID()
	post.AuthorUsername = c.GetString("username")

	if err := post.SavePost(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post created and uploaded successfully",
	})
}

func UpdatePost(c *gin.Context) {
	postID := c.Param("id")

	var updateFields map[string]interface{}
	if err := c.BindJSON(&updateFields); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ID, err := primitive.ObjectIDFromHex(postID)
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

	c.JSON(200, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	var postID string
	if err := c.ShouldBindJSON(&postID); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "post deleted successfully",
	})
}

func GetBatchOfPosts(c *gin.Context) {
	batchNumber, err := strconv.Atoi(c.Param("batchNumber"))
	if err != nil || batchNumber < 1 {
		c.JSON(400, gin.H{"error": "invalid batch number"})
		return
	}

	posts, err := GetPostsBatch(batchNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": posts,
	})
}

func GetPostByID(c *gin.Context) {
	postID := c.Param("postID")
	ID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalig postID"})
		return
	}
	post, err := GetPostByIDDB(ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": post,
	})
}

// +-------------------+
// |   POST COMMENTS   |
// +------------------+

func CreateComment(c *gin.Context) {
	var comment = NewComment()
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	comment.AuthorUsername = c.GetString("username")
	comment.ID = primitive.NewObjectID()

	if err := comment.SaveComment(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post commented successfully",
	})
}

func UpdateComment(c *gin.Context) {
	var comment = NewComment()
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := comment.UpdateComment(); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Post commented successfully",
	})
}

func DeleteComment(c *gin.Context) {
	commentID := c.Param("commentID")
	ID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalig postID"})
		return
	}
	post, err := GetPostByIDDB(ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": post,
	})
}

// +--------------------+
// |   POST + ALBUMS   |
// +-------------------+

func AddToAlbum(c *gin.Context) {}

func DeleteFromAlbum(c *gin.Context) {}
