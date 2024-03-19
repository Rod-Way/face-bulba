package posts

import (
	"context"
	db "faceBulba/database"
	u "faceBulba/internal/user"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID             primitive.ObjectID `json:"-" bson:"_id"`
	AuthorUsername string             `json:"-" bson:"author"`
	AlbumsIDs      []string           `json:"-" bson:"albums_ids"`
	Text           string             `json:"text" bson:"text"`
	FilesURL       []string           `json:"files_url" bson:"files_url"`
	Tags           []string           `json:"tags" bson:"tags"`
	Comments       []Comment          `json:"-" bson:"comments"`
	IsUpdated      bool               `json:"-" bson:"is_updated"`
	CreatedAt      time.Time          `json:"-" bson:"createdAt"`
}

type Comment struct {
	ID             primitive.ObjectID `json:"-"`
	PostID         primitive.ObjectID `json:"post_id"`
	AuthorUsername string             `json:"-" bson:"author"`
	Text           string             `json:"text"`
	CreatedAt      time.Time          `json:"-"`
}

func NewPost() *Post {
	return new(Post)
}

func NewComment() *Comment {
	return new(Comment)
}

func (p *Post) SavePost() error {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	if _, err = collection.InsertOne(context.TODO(), p); err != nil {
		return fmt.Errorf("database inserting error: %v", err)
	}

	if err = addPostToUser(p.AuthorUsername, p.ID); err != nil {
		return fmt.Errorf("database adding post to user error: %v", err)
	}

	return nil
}

func GetField(postID primitive.ObjectID, fieldName string) (interface{}, error) {
	client, col, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	var post Post
	err = col.FindOne(ctx, bson.M{"_id": postID}).Decode(&post)
	if err != nil {
		return nil, fmt.Errorf("failed to find post: %v", err)
	}

	switch fieldName {
	case "author":
		return post.AuthorUsername, nil
	case "albums_ids":
		return post.AlbumsIDs, nil
	case "text":
		return post.Text, nil
	case "files_url":
		return post.FilesURL, nil
	case "tags":
		return post.Tags, nil
	case "comments":
		return post.Comments, nil
	case "is_updated":
		return post.IsUpdated, nil
	case "createdAt":
		return post.CreatedAt, nil
	default:
		return nil, fmt.Errorf("can not get this field: %s", fieldName)
	}
}

func (p *Post) UpdatePost(newPostData *Post) error {
	client, col, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"_id": p.ID}
	updateResult, err := col.UpdateOne(ctx, filter, bson.M{"$set": newPostData})
	if err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}

	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}

func GetPostByIDDB(postID primitive.ObjectID) (*Post, error) {
	client, col, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	var post Post
	err = col.FindOne(ctx, bson.M{"_id": postID}).Decode(&post)
	if err != nil {
		return nil, fmt.Errorf("failed to find post: %v", err)
	}

	return &post, nil
}

// Getting batch for "Endless" feed
func GetPostsBatch(batchNumber int) ([]Post, error) {
	client, col, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	skip := (batchNumber - 1) * 25

	cursor, err := col.Find(ctx, bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(25))
	if err != nil {
		return nil, fmt.Errorf("failed to query MongoDB collection: %v", err)
	}
	defer cursor.Close(ctx)

	var posts []Post
	err = cursor.All(ctx, &posts)
	if err != nil {
		return nil, fmt.Errorf("failed to decode posts: %v", err)
	}

	return posts, nil
}

func DeletePostByID(postID primitive.ObjectID) error {
	client, col, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	filter := bson.M{"_id": postID}
	deleteResult, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("post not found")
	}

	username, err := GetField(postID, "username")
	if err != nil {
		return fmt.Errorf("failed to find username: %v", err)
	}
	un, ok := username.(string)
	if !ok {
		return fmt.Errorf("invalid username: %v", err)
	}
	err = removePostFromUser(un, postID)
	if err != nil {
		return fmt.Errorf("failed to remove post from user: %v", err)
	}

	return nil
}

func addPostToUser(username string, postID primitive.ObjectID) error {
	client, col, ctx, cancel, err := db.GetDB("users")
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	filter := bson.M{"username": username}
	var user u.User
	err = col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	ID := postID.Hex()

	user.Posts = append(user.Posts, ID)

	update := bson.M{"$set": bson.M{"posts": user.Posts}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func removePostFromUser(username string, postID primitive.ObjectID) error {
	client, col, ctx, cancel, err := db.GetDB("users")
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	filter := bson.M{"username": username}
	var user u.User
	err = col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	postIDString := postID.Hex()

	index := -1
	for i, id := range user.Posts {
		if id == postIDString {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("postID not found in user's posts")
	}

	user.Posts = append(user.Posts[:index], user.Posts[index+1:]...)

	update := bson.M{"$set": bson.M{"posts": user.Posts}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}
