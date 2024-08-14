package controllers

import (
	"net/http"
	domain "task/Domain"
	infrastructure "task/Infrastructure"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(usecase domain.UserUsecase) *UserController {
	return &UserController{
		UserUsecase: usecase,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {

	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = uc.UserUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user registerd sucessfully", "Id": user.ID})
}
func (uc *UserController) DeleteUser(c *gin.Context) {
	claims := c.MustGet("claims").(*domain.Claims)
	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
		return
	}
	id := c.Param("id")
	err := uc.UserUsecase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "claims.Role": claims.Role})

}
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	claims := c.MustGet("claims").(*domain.Claims)

	if claims.Role != "admin" && claims.Id != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
		return
	}
	user, err := uc.UserUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (uc *UserController) GetUsers(c *gin.Context) {
	claims := c.MustGet("claims").(*domain.Claims)

	if claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
		return
	}
	users, err := uc.UserUsecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (uc *UserController) UpdateUser(c *gin.Context) {
	claims := c.MustGet("claims").(*domain.Claims)
	id := c.Param("id")

	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if claims.Role != "admin" && claims.UserID != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are unauthorized"})
		return
	}
	NewId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.ID = NewId
	err = uc.UserUsecase.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
func (uc *UserController) LoginUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// TODO: Implement user login logic
	// storedUser, ok := data.FindOneUser(data.Users, bson.M{"email": user.Email})
	storedUser, err := uc.UserUsecase.Login(user)
	if err == mongo.ErrNoDocuments {
		c.JSON(500, gin.H{"error": "User not found"})
		return
	}
	if storedUser.Email == "" || !infrastructure.ComparePassword(user.Password, storedUser.Password) {
		c.JSON(401, gin.H{"error": "Invalid email or password "})
		return
	}

	claims := &domain.Claims{
		UserID: storedUser.ID.Hex(),
		Email:  storedUser.Email,
		Role:   storedUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	tokenString, err := uc.UserUsecase.GeneratesToken(*claims)

	if err != nil {
		c.JSON(500, gin.H{"error": "error while generating jwt token", "err": err})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": tokenString, "Id": claims.UserID})

}
