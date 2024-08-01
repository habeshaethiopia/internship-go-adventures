package controllers

import (
	"net/http"

	"task/data"
	"task/models"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Tasks)
}
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	for _, t := range data.Tasks {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, val := range data.Tasks {
		if val.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			data.SaveTasksToFile(data.Tasks, "tasks.json")
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
func PostTask(ctx *gin.Context) {
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if newTask.ID == "" || newTask.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID and Name are required"})
		return
	}
	
	data.Tasks = append(data.Tasks, newTask)
	data.SaveTasksToFile(data.Tasks, "tasks.json")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}
func PutTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range data.Tasks {
		if task.ID == id {
			// Update only the specified fields
			if updatedTask.Title != "" {
				data.Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				data.Tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				data.Tasks[i].Status = updatedTask.Status
			}
			data.SaveTasksToFile(data.Tasks, "tasks.json")
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
