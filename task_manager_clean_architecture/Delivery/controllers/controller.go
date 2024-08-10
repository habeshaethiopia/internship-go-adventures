package controllers

import (
	"fmt"
	"net/http"
	domain "task/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	claims := c.MustGet("claims").(*domain.Claims)
	var task domain.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.UserID = userID
	err = tc.TaskUsecase.CreateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}
func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := tc.TaskUsecase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task domain.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = tc.TaskUsecase.UpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

type UserController struct {
	UserUsecase domain.UserUsecase
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
	c.JSON(http.StatusCreated, user)
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
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.UserUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
func (uc *UserController) GetUsers(c *gin.Context) {
	claims := c.MustGet("claims").(*domain.Claims)

	fmt.Println(claims, "claims")
	d := c.Param("id")

	if claims.Role != "admin" && claims.UserID != d {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized", "claims": claims})
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
	d := c.Param("id")

	if claims.Role != "admin" && claims.Id != d {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are anauthorized"})
		return
	}
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		c.JSON(500, gin.H{"error": "error while checking user"})
		return
	}
	if storedUser.Email == "" || bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password " + storedUser.Password})
		return
	}
	var jwtsecret = []byte("btavg5IIbDkyYmUWjN6C")
	claims := &domain.Claims{
		UserID: storedUser.ID.Hex(),
		Email:  storedUser.Email,
		Role:   storedUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtsecret)
	color.Green("secret: %s\n", jwtsecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "error while generating jwt token", "err": err})
		return
	}

	c.JSON(200, gin.H{"message": "User logged in successfully", "token": tokenString, "user": storedUser, "secret": jwtsecret})

}
