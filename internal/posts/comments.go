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

	var post struct {
		Comments []Comment `bson:"comments"`
	}
	err = collection.FindOne(ctx, bson.M{"comments.id": commentID}).Decode(&post)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("comment not found with ID %s", commentID.Hex())
		}
		return nil, fmt.Errorf("failed to decode post: %v", err)
	}

	for _, comment := range post.Comments {
		if comment.ID == commentID {
			switch fieldName {
			case "id":
				return comment.ID, nil
			case "post_id":
				return comment.PostID, nil
			case "author":
				return comment.AuthorUsername, nil
			case "text":
				return comment.Text, nil
			case "CreatedAt":
				return comment.CreatedAt, nil
			default:
				return nil, fmt.Errorf("can not get this field: %s", fieldName)
			}
		}
	}

	return nil, fmt.Errorf("comment not found with ID %s", commentID.Hex())
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
