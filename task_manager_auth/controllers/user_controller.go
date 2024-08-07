package controllers

import (
	// "context"
	"fmt"
	"os"
	"time"

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
	claims := c.MustGet("claims").(*models.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
		return
	}
	c.IndentedJSON(http.StatusOK, data.FindAllUsers(data.Users))
}
func GetUserById(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
		return
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	fmt.Printf("id: %v\n", id)
	result, err := data.FindOneUser(data.Users, bson.M{"_id": id})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
func UpdateUser(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)
	d := c.Param("id")
	id, err := primitive.ObjectIDFromHex(d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if claims.Role != "admin" && claims.UserID != d {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedUser.ID == primitive.NilObjectID {
		updatedUser.ID = id
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in body must match ID in URL", "The ID": updatedUser.ID})
		return
	}
	result, er := data.FindOneUser(data.Users, bson.M{"_id": id})
	if er != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if updatedUser.Email == "" {
		updatedUser.Email = result.Email
	}
	if updatedUser.Password == "" {
		updatedUser.Password = result.Password
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "error while hashing the password"})
			return
		}
		updatedUser.Password = string(hashedPassword)
	}
	if updatedUser.Role == "" {
		updatedUser.Role = result.Role
	} else {
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
			return
		}
	}
	updateResult, err := data.UpdateUser(data.Users, bson.M{"_id": id}, bson.M{"$set": updatedUser})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if updateResult.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)

}

func RegisterUser(c *gin.Context) {

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return

	}
	//TODO: Impliment user registration logic

	exist, err := data.FindOneUser(data.Users, bson.M{"email": user.Email})
	color.Cyan("user: %v\n", exist)

	if err != nil {
		c.JSON(500, gin.H{"error": "error while checking user", "err": err})
		return
	}
	if exist.Email == user.Email {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while hashing the password"})
		return
	}
	user.ID = primitive.NewObjectID()
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
	storedUser, err := data.FindOneUser(data.Users, bson.M{"email": user.Email})
	if err == mongo.ErrNoDocuments {
		c.JSON(500, gin.H{"error": "error while checking user"})
		return
	}

	if storedUser.Email == "" || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password", "storedUser": storedUser})
		return
	}
	var jwtsecret = []byte(os.Getenv("JWT_SECRET_KEY"))
	claims := &models.Claims{
		UserID: storedUser.ID.Hex(),
		Email:  storedUser.Email,
		Role:   storedUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtsecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while generating jwt token", "err": err})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": tokenString, "user": storedUser})
}

func DeleteUser(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized", "role": claims.Role})
		return
	}
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	result, err := data.DeleteUser(data.Users, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}