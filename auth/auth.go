package auth

import (
	"errors"
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// This function generates a new JSON Web Token, also accepts the users username as a input
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user"] = username
	claims["aud"] = "go-social.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	// The token will expire after 24 hours / 1 day
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		err = errors.New("something went wrong")
		log.Fatal("Auth error:", err)
		return "", err
	}
	return newToken, nil
}

// This function checks if the JWT is still valid
func CheckJWT(token string) (string, error) {
	var username string

	t, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return username, err
	} else if !t.Valid {
		err = errors.New("invalid token")
		log.Fatal("Auth error:", err)
		return username, err
	} else if t.Claims.(jwt.MapClaims)["aud"] != "go-social.jwtgo.io" {
		err = errors.New("invalid aud")
		log.Fatal("Auth error:", err)
		return username, err
	} else if t.Claims.(jwt.MapClaims)["iss"] != "jwtgo.io" {
		err = errors.New("invalid iss")
		log.Fatal("Auth error:", err)
		return username, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		username = claims["user"].(string)
	}
	return username, nil
}
