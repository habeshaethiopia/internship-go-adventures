package controllers

import (
	"net/http"
	domain "task/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (tc *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")

	claims := ctx.MustGet("claims").(*domain.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, er := tc.TaskUsecase.GetTaskByID(id)
	if er != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	if claims.Role != "admin" && result.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err = tc.TaskUsecase.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	claims := c.MustGet("claims").(*domain.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task, err := tc.TaskUsecase.GetTaskByID(id)

	if err != nil || task.UserID != userID {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Task not found"})
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
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No tasks found"})
		return
	}
	claims := c.MustGet("claims").(*domain.Claims)
	userID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID","clamis": claims.UserID})
		return
	}
	if claims.Role == "admin" {
		c.IndentedJSON(http.StatusOK, tasks)
		return
	}
	
	// Find all tasks associated with the logged user's ID
	var taskList []domain.Task
	for _, task_s := range tasks {
		if task_s.UserID == userID {
			taskList = append(taskList, *task_s)
		}
	}

	

	c.IndentedJSON(http.StatusOK, taskList)
	
}
func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task domain.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	task.ID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tc.TaskUsecase.UpdateTask(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}
