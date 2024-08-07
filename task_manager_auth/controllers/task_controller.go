package controllers

import (
	// "context"
	"fmt"
	// "image/color"
	// "image/color"
	"net/http"

	"task/data"
	"task/models"

	// "task/models"
	// "fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetTasks(c *gin.Context) {
	// Get the logged user's ID
	claims := c.MustGet("claims").(*models.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID","clamis": claims.UserID})
		return
	}
	result := data.FindAll(data.Tasks)
	if claims.Role == "admin" {
		c.IndentedJSON(http.StatusOK, result)
		return
	}
	
	// Find all tasks associated with the logged user's ID
	var taskList []models.Task
	for _, task_s := range result {
		if task_s.UserID == userID {
			taskList = append(taskList, *task_s)
		}
	}

	

	c.IndentedJSON(http.StatusOK, taskList)
}
func GetTaskById(c *gin.Context) {
	claims := c.MustGet("claims").(*models.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	id, e := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil || e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if claims.Role != "admin" {
		result, err := data.FindOne(data.Tasks, bson.M{"_id": id})
		if err != nil || err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, result)
	} else {
		result:= data.FindAll(data.Tasks)
		
		for _, task := range result {
			if task.UserID == userID && task.ID == id {
				c.IndentedJSON(http.StatusOK, task)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
	}
		
		
}
func DeleteTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	claims := ctx.MustGet("claims").(*models.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, er := data.FindOne(data.Tasks, bson.M{"_id": id})
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	fmt.Println(result)
	if claims.Role != "admin" && result.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	

	err = data.Delete(data.Tasks, bson.M{"_id": id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})

}
func PostTask(ctx *gin.Context) {
	var newTask models.Task
	claims := ctx.MustGet("claims").(*models.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	newTask.UserID = userID

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newTask.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	newTask.ID = primitive.NewObjectID()
	// Check if the ID is already in use
	id := newTask.ID
	existingTask, err := data.FindOne(data.Tasks, bson.M{"_id": id})
	fmt.Println(id, err, mongo.ErrNoDocuments)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check ID"})
			return
		}
	} else if existingTask.ID != primitive.NilObjectID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID already in use"})
		return
	}

	err = data.InsertOne(data.Tasks, newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert task"})
		return
	}

	ctx.JSON(http.StatusOK, newTask)
}

func PutTask(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*models.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	d := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(d)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	result, er := data.FindOne(data.Tasks, bson.M{"_id": id})
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	fmt.Println(result)
	if claims.Role != "admin" && result.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("updatedTask.id: %s   id: %s\n", updatedTask.ID, id)
	if updatedTask.ID == primitive.NilObjectID {
		updatedTask.ID = id
		color.Green("updatedTask.ID: %s", updatedTask.ID)
	}
	if updatedTask.ID != id {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID in body must match ID in URL"})
		return
	}
	if updatedTask.Title == "" {
		updatedTask.Title = result.Title
	}
	if updatedTask.Description == "" {
		updatedTask.Description = result.Description
	}
	if updatedTask.DueDate.IsZero() {
		updatedTask.DueDate = result.DueDate
	}
	if updatedTask.Status == "" {
		updatedTask.Status = result.Status
	}
	if updatedTask.UserID == primitive.NilObjectID {
		updatedTask.UserID = result.UserID
	}
	color.Red("updatedTask: %v", updatedTask)
	updateResult, err := data.UpdateOne(data.Tasks, bson.M{"_id": id}, bson.M{"$set": updatedTask})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	if updateResult.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, updatedTask)
}
