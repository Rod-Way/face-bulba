package albums

import (
	"log"
	"time"

	db "faceBulba/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	Users     []string           `json:"-" bson:"users"`
	Name      string             `json:"name" bson:"name"`
	Content   map[string]string  `json:"content" bson:"content"`
	Tags      []string           `json:"tags" bson:"tags"`
	CreatedAt time.Time          `json:"-" bson:"created_at"`
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

func UpdateAlbum() {}

func DeleteAlbum() {}
