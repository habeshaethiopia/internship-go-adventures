package main

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
var users = make(map[string]*User)
var jwtSecret = []byte("btavg5IIbDkyYmUWjN6C")


func mainc() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Go Authentication and Authorization tutorial!",
		})
	})

	router.POST("/register", registerUser)
	router.POST("/login", loginUser)
	router.GET("/secure", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a secure route"})
	  })

	router.Run()
}
func registerUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return

	}
	//TODO: Impliment user registration logic
	if _, ok := users[user.Email]; ok {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while hashing the password"})
		return
	}
	user.Password = string(hashedPassword)
	users[user.Email] = &user

	c.JSON(200, gin.H{"message": "user registerd successfully"})
}
func loginUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// TODO: Implement user login logic
	storedUser, ok := users[user.Email]
	if !ok || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}
	var jwtsecret = []byte("btavg5IIbDkyYmUWjN6C")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"email":   user.Email,
	})
	tokenString, err := token.SignedString(jwtsecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while generating jwt token"})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": tokenString})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT validation logic
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return

		}

		c.Next()
	}
}
