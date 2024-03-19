package user

import (
	"errors"
	"log"
	"time"

	CO "faceBulba/config"
	db "faceBulba/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID              primitive.ObjectID `json:"-" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Surname         string             `json:"surname" bson:"surname"`
	Username        string             `json:"username" bson:"username"`
	Email           string             `json:"email" bson:"email"`
	Posts           []string           `json:"-" bson:"posts"`   // All user post IDs
	Albums          []string           `json:"-" bson:"albums"`  // All user albums IDs
	Friends         []string           `json:"-" bson:"friends"` // All users friends
	EmailVerifiedAt string             `json:"-" bson:"email_verifaed_at,omitempty"`
	Password        string             `json:"password,omitempty" bson:"password"`
	CreatedAt       string             `json:"-" bson:"created_at"`
}

// Structure for login
type LogInUser struct {
	User     string `json:"user"`                     // Username or email
	Password string `json:"password" bson:"password"` // Password
}

func NewUser() *User {
	return new(User)
}
func NewLogInUser() *LogInUser {
	return new(LogInUser)
}

// Saves a new user to the database
func (u *User) SaveUser() error {
	client, collection, ctx, cancel, err := db.GetDB("users")
	if err != nil {
		log.Fatal("Failed to get MongoDB client and collection:", err)
		return err
	}
	defer client.Disconnect(ctx)
	defer cancel()

	// Check if the user already exists
	var existingUser User
	err = collection.FindOne(ctx, bson.M{"$or": []bson.M{{"username": u.Username}, {"email": u.Email}}}).Decode(&existingUser)
	if err == nil {
		errr := errors.New("User already exists")
		log.Print("MongoDB error:", errr)
		return errr
	} else if err != mongo.ErrNoDocuments {
		log.Fatal("MongoDB error:", err)
		return err
	}

	// Hash the user's password
	hashedPassword, err := CO.HashPassword(u.Password)
	if err != nil {
		log.Fatal("MongoDB hash error:", err)
		return err
	}

	// Add the new user
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now().Format("2006-12-30 15:04:05")
	u.EmailVerifiedAt = ""
	u.Posts = nil
	u.Friends = nil
	u.Albums = nil

	_, err = collection.InsertOne(ctx, u)
	return err
}

/*
Login a user
If everything goes well then return the users username along with a nil error
The username will be used to generate a new JWT Token
A user should be able to login via their unique username or as an alternative
they can also user their unique email address to login
*/
func (u *LogInUser) UserLogin() (string, error) {
	client, collection, ctx, cancel, err := db.GetDB("users")
	if err != nil {
		log.Fatal("Failed to get MongoDB client and collection:", err)
		return "", err
	}
	defer client.Disconnect(ctx)
	defer cancel()

	var user User
	err = collection.FindOne(ctx, bson.M{"$or": []bson.M{{"username": u.User}, {"email": u.User}}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		errr := errors.New("the user does not exist")
		log.Print("Error:", errr)
		return "", errr
	} else if err != nil {
		return "", err
	}

	// Check password
	err = CO.CheckPassword(user.Password, u.Password)
	if err != nil {
		errr := errors.New("wrong password")
		log.Print("Error:", errr, err)
		return "", errr
	}

	return user.Username, nil
}
