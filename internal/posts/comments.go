package posts

import (
	db "faceBulba/database"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (com *Comment) SaveComment() error {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"_id": com.PostID}
	var post Post
	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return fmt.Errorf("failed to find post: %v", err)
	}

	post.Comments = append(post.Comments, *com)

	update := bson.M{"$set": bson.M{"comments": post.Comments}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update post: %v", err)
	}

	return nil
}

func (com *Comment) UpdateComment() error {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"_id": com.ID, "comments._id": com.PostID}

	update := bson.M{"$set": bson.M{"comments.$.text": com.Text}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update comment: %v", err)
	}

	return nil
}

func DeleteCommentF(commentID primitive.ObjectID) error {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"comments._id": commentID}

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}
