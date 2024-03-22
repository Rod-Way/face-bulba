package albums

import (
	"errors"
	"fmt"
	"log"

	db "faceBulba/database"
	u "faceBulba/internal/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Album struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	Users     []string           `json:"-" bson:"users"`
	Name      string             `json:"name" bson:"name"`
	Content   map[string]string  `json:"content" bson:"content"`
	Tags      []string           `json:"tags" bson:"tags"`
	CreatedAt string             `json:"-" bson:"created_at"`
}

func NewAlbum() *Album {
	return new(Album)
}

func (a *Album) SaveAlbum() error {
	client, collection, ctx, cancel, err := db.GetDB("albums")
	if err != nil {
		log.Fatal("Failed to get MongoDB client and collection:", err)
		return err
	}
	defer client.Disconnect(ctx)
	defer cancel()

	_, err = collection.InsertOne(ctx, a)
	return err
}

func GetField(albumID primitive.ObjectID, fieldName string) (interface{}, error) {
	client, col, ctx, cancel, err := db.GetDB("albums")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	var album Album
	err = col.FindOne(ctx, bson.M{"_id": albumID}).Decode(&album)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("post not found with ID %s", albumID.Hex())
		}
		return nil, fmt.Errorf("failed to decode post: %v", err)
	}

	var fieldValue interface{}
	switch fieldName {
	case "_id":
		fieldValue = album.ID
	case "users":
		fieldValue = album.Users
	case "name":
		fieldValue = album.Name
	case "content":
		fieldValue = album.Content
	case "tags":
		fieldValue = album.Tags
	case "created_at":
		fieldValue = album.CreatedAt
	default:
		return nil, fmt.Errorf("can not get this field: %s", fieldName)
	}

	if fieldValue == nil {
		return nil, fmt.Errorf("field %s does not exist in post with ID %s", fieldName, albumID.Hex())
	}

	return fieldValue, nil
}

func FindUser(albumID primitive.ObjectID, username string) (string, error) {
	client, col, ctx, cancel, err := db.GetDB("albums")
	if err != nil {
		return "", fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	var album struct {
		Users []string `bson:"users"`
	}
	err = col.FindOne(ctx, bson.M{"_id": albumID, "users": bson.M{"$in": []string{username}}}).Decode(&album)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", fmt.Errorf("user '%s' is not an author of the album", username)
		}
		return "", fmt.Errorf("failed to decode album: %v", err)
	}

	return username, nil
}

func isAllowedField(field string) bool {
	allowedFields := map[string]bool{
		"users":      true,
		"name":       true,
		"content":    true,
		"tags":       true,
		"created_at": true,
	}
	return allowedFields[field]
}

func UpdateField(albumID primitive.ObjectID, field string, value interface{}) error {
	client, col, ctx, cancel, err := db.GetDB("albums")
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)
	defer cancel()

	filter := bson.M{"_id": albumID}

	update := bson.M{"$set": bson.M{field: value}}

	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAlbumByID(albumID primitive.ObjectID) error {
	client, col, ctx, cancel, err := db.GetDB("albums")
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	filter := bson.M{"_id": albumID}
	deleteResult, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}

func GetAlbumByIDDB(albumID primitive.ObjectID) (*Album, error) {
	client, col, ctx, cancel, err := db.GetDB("albums")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)
	defer cancel()

	var album Album
	err = col.FindOne(ctx, bson.M{"_id": albumID}).Decode(&album)
	if err != nil {
		return nil, fmt.Errorf("failed to find post: %v", err)
	}

	return &album, nil
}

func AddAlbumToUser(albumID primitive.ObjectID, username string) error {
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

	ID := albumID.Hex()

	user.Albums = append(user.Albums, ID)

	update := bson.M{"$set": bson.M{"albums": user.Albums}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func DeleteAlbumFromUser(albumID primitive.ObjectID, username string) error {
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

	albumIDString := albumID.Hex()

	index := -1
	for i, id := range user.Albums {
		if id == albumIDString {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("albumID not found in user's albums")
	}

	user.Albums = append(user.Albums[:index], user.Albums[index+1:]...)

	update := bson.M{"$set": bson.M{"albums": user.Albums}}
	_, err = col.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}
