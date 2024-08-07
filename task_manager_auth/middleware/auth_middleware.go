package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"os"

	"task/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtsecret = []byte("btavg5IIbDkyYmUWjN6C")

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

		tokenString := authParts[1]
        claims := &models.Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtsecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid token", "err": err, "token": token})
            c.Abort()
            return
        }
		
        c.Set("claims", claims)
		color.Green("claims: %v\n",claims)
		userID, err := primitive.ObjectIDFromHex(claims.UserID)
		if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID","clamis": claims, "claims.Id": userID})
		return
	}
        c.Next()
	}
}