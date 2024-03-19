package posts

import (
	"strconv"
	"time"

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
	post.CreatedAt = time.Now()

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

	if err := DeletePostByID(objID); err != nil {
		c.JSON(500, gin.H{
			"error": "Can not delete",
		})
	}

	if err := removePostFromUser(auth, objID); err != nil {
		c.JSON(500, gin.H{
			"error": "Can not delete",
		})
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
	comment.CreatedAt = time.Now()

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
	commentID := c.Param("id")

	var newText string
	if err := c.ShouldBindJSON(&newText); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid commentID",
		})
		return
	}

	if err := UpdateCommentF(ID, newText); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Comment updated successfully",
	})
}

func DeleteComment(c *gin.Context) {
	var commentID string
	if err := c.ShouldBindJSON(&commentID); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "invalid postID",
		})
	}

	// Check author and user
	author, err := GetCommentField(objID, "author")
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

	if err := DeleteCommentF(objID); err != nil {
		c.JSON(500, gin.H{
			"error": "Can not delete",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Comment deleted successfully",
	})
}

// +--------------------+
// |   POST + ALBUMS   |
// +-------------------+

func AddToAlbum(c *gin.Context) {}

func DeleteFromAlbum(c *gin.Context) {}
