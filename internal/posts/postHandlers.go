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
	var post = NewPost()
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	author, err := GetField(post.ID, "username")
	if err != nil || author != c.GetString("username") {
		c.JSON(400, gin.H{
			"error": "invalig author",
		})
		return
	}

	if err := post.UpdatePost(post); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "post updated successfully",
	})
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

	// if err := post.SavePost(); err != nil {
	// 	c.JSON(500, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(200, gin.H{
		"message": "Post commented successfully",
	})
}

func UpdateComment(c *gin.Context) {}

func DeleteComment(c *gin.Context) {}

// +--------------------+
// |   POST + ALBUMS   |
// +-------------------+

func AddToAlbum(c *gin.Context) {}

func DeleteFromAlbum(c *gin.Context) {}
