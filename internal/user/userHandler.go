package user

import (
	"strconv"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	"github.com/lsgser/go-social/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// This function will handle data from the add_user endpoint
func AddUser(c *gin.Context) {
	user := NewUser()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Check if the required strings are not empty
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Surname) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}
	/*
	     Make the input strings more presentable before storing it to the database
	     The email should contain only lowercase characters i.e if a user
	     inputs the email as EXample@something.com we'll format it to
	     example@something.com
	   The name and surname/lastname should contain the first letter as a Uppercase
	     if the user enters their name as john we'll format it to John
	*/

	// Creating a Header Form Instance
	upper := cases.Title(language.English)

	user.ID = primitive.NewObjectID()

	user.Name = upper.String(strings.TrimSpace(strings.ToLower(user.Name)))
	user.Surname = upper.String(strings.TrimSpace(strings.ToLower(user.Surname)))
	user.Email = strings.ToLower(user.Email)
	//Check if the email format is valid
	if checkmail.ValidateFormat(user.Email) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email format",
		})
		return
	}
	/*
	   Save the user to the database using the SaveUser method
	   that we created in the user.go file
	*/
	err := user.SaveUser()
	if err != nil {
		if err.Error() == "User already exists" {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		//Other errors from the SaveUser() method are internal server errors
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}

func LoginUser(c *gin.Context) {
	user_login := NewLogInUser()
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Check if the required strings are not empty
	if strings.TrimSpace(user_login.User) == "" && strings.TrimSpace(user_login.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}
	//Validate the credentials
	username, err := user_login.UserLogin()
	if err != nil {
		if err.Error() == "The user does not exist" || err.Error() == "Wrong password" {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	//Generate the token
	token, err := auth.GenerateJWT(username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func CheckUser(c *gin.Context) {
	token := c.Param("token")
	if strings.TrimSpace(token) == "" {
		c.JSON(400, gin.H{
			"error": "Token not provided",
		})
		return
	}
	err := auth.CheckJWT(token)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"isAuthenticated": true,
	})
}

func GetBachOfUSers(c *gin.Context) {
	batchNumber, err := strconv.Atoi(c.Param("batchNumber"))
	if err != nil || batchNumber < 1 {
		c.JSON(400, gin.H{"error": "invalid batch number"})
		return
	}

	posts, err := GetUsersBatch(batchNumber)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": posts,
	})
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	ID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalig postID"})
		return
	}
	post, err := GetUserByIDDB(ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"response": post,
	})
}
