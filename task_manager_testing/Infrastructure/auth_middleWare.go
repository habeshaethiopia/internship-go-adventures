package infrastructure

import (
	"fmt"
	"net/http"
	"strings"
	domain "task/Domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware(secret []byte) gin.HandlerFunc {
	
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
		claims := &domain.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token", "err": err, "token": token, "secret": secret})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		color.Green("claims: %v\n", claims)
		userID, err := primitive.ObjectIDFromHex(claims.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID", "clamis": claims, "claims.Id": userID})
			return
		}
		c.Next()
	}
}
