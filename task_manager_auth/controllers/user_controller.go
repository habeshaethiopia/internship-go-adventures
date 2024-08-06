package controllers

import (
	// "context"
	"fmt"
	"os"
	// "image/color"
	// "image/color"
	"net/http"

	"task/data"
	"task/models"

	// "task/models"
	// "fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.FindAll(data.Tasks))
}
func GetUserById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	fmt.Printf("id: %v\n", id)
	result, err := data.FindOne(data.Tasks, bson.M{"_id": id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}


func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return

	}
	//TODO: Impliment user registration logic
	
	exist,err := data.FindOneUser(data.Users, bson.M{"email":user.Email})
	color.Cyan("user: %v\n", exist)
	
	if err !=  nil {
		c.JSON(500, gin.H{"error": "error while checking user" ,"err":err})
		return
	}
	if exist.Email ==  user.Email{
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}
	
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while hashing the password"})
		return
	}
	user.Password = string(hashedPassword)
	data.InsertUser(data.Users, user)

	c.JSON(200, gin.H{"message": "user registerd successfully"})
}
func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// TODO: Implement user login logic
	// storedUser, ok := data.FindOneUser(data.Users, bson.M{"email": user.Email})
	storedUser,err := data.FindOneUser(data.Users, bson.M{"email":user.Email})
	if err ==  mongo.ErrNoDocuments {
		c.JSON(500, gin.H{"error": "error while checking user"})
		return
	}
	
	if  storedUser.Email ==  "" || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password" ,"storedUser":storedUser})
		return
	}
var jwtsecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	// var jwtsecret = os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"email":   user.Email,
		"role":    user.Role,
	})
	tokenString, err := token.SignedString(jwtsecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while generating jwt token", "err": err})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": tokenString, "user": storedUser})
}