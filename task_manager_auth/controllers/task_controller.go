package controllers

import (
	// "context"
	"fmt"
	// "image/color"
	"net/http"

	"task/data"
	"task/models"

	// "task/models"
	// "fmt"
	// "github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.FindAll(data.Tasks))
}
func GetTaskById(c *gin.Context) {
    id,err:= primitive.ObjectIDFromHex(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
    }
	fmt.Printf("id: %v\n",id)
    result, err := data.FindOne(data.Tasks, bson.M{"_id": id})
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
        return
    }

    c.IndentedJSON(http.StatusOK, result)
}

func DeleteTask(ctx *gin.Context) {
	id,err:= primitive.ObjectIDFromHex(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
    }
    

	result,er:=data.FindOne(data.Tasks, bson.M{"_id": id})
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	fmt.Println(result)

	err = data.Delete(data.Tasks, bson.M{"_id": id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})

}
func PostTask(ctx *gin.Context) {
    var newTask models.Task
	
	
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
	fmt.Println( id, err, mongo.ErrNoDocuments)
    if err != nil {
        if err!= mongo.ErrNoDocuments {
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
    d := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(d)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	result,er:=data.FindOne(data.Tasks, bson.M{"_id": id})
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	fmt.Println(result)

    var updatedTask models.Task

    if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	fmt.Printf("updatedTask.id: %s   id: %s\n", updatedTask.ID, id)
	if updatedTask.ID == primitive.NilObjectID{
		updatedTask.ID=id
	}
	if updatedTask.ID != id {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID in body must match ID in URL"})
		return
	}
	if updatedTask.Title == "" {
		updatedTask.Title=result.Title
	}
	if updatedTask.Description == "" {
		updatedTask.Description=result.Description
	}
	if updatedTask.DueDate.IsZero() {
		updatedTask.DueDate=result.DueDate
	}

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