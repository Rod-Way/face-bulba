package posts

import (
	"errors"
	db "faceBulba/database"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func UpdateCommentF(commentID primitive.ObjectID, value string) error {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"comments.id": commentID}

	update := bson.M{"$set": bson.M{"comments.$.text": value}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update comment: %v", err)
	}

	return nil
}

func GetCommentField(commentID primitive.ObjectID, fieldName string) (interface{}, error) {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	var comm Comment
	err = collection.FindOne(ctx, bson.M{"comments.id": commentID}).Decode(&comm)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("post not found with ID %s", commentID.Hex())
		}
		return nil, fmt.Errorf("failed to decode post: %v", err)
	}

	var fieldValue interface{}
	switch fieldName {
	case "id":
		fieldValue = comm.ID
	case "post_id":
		fieldValue = comm.PostID
	case "author":
		fieldValue = comm.AuthorUsername
	case "text":
		fieldValue = comm.Text
	case "CreatedAt":
		fieldValue = comm.CreatedAt
	default:
		return nil, fmt.Errorf("can not get this field: %s", fieldName)
	}

	// if fieldValue == nil {
	// 	return nil, fmt.Errorf("can not get this field: %s", fieldName)
	// }

	return fieldValue, nil
}

func DeleteCommentF(commentID primitive.ObjectID) error {
	client, collection, ctx, cancel, err := db.GetDB("posts")
	if err != nil {
		return fmt.Errorf("database error: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"comments.id": commentID}

	deleteResult, err := collection.UpdateOne(ctx, filter, bson.M{"$pull": bson.M{"comments": bson.M{"id": commentID}}})
	if err != nil {
		return fmt.Errorf("failed to delete comment: %v", err)
	}

	if deleteResult.ModifiedCount == 0 {
		return fmt.Errorf("comment not found")
	}

	return nil
}
