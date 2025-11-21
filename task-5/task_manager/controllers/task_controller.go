// package data

// import (
// 	"errors"
// 	"task_manager/models"
// )

// func GetAllTasks() []models.Task {
// 	return models.Tasks
// }

// func GetTaskById(id string) (models.Task, error) {
// 	for _, t := range models.Tasks {
// 		if t.ID == id {
// 			return t, nil
// 		}
// 	}
// 	return models.Task{}, errors.New("task not found")
// }

// func AddNewTask(task models.Task) models.Task {
// 	models.Tasks = append(models.Tasks, task)
// 	return task
// }

// func UpdateTask(id string, updated models.Task) (models.Task, error) {
// 	for i, t := range models.Tasks {
// 		if t.ID == id {
// 			models.Tasks[i] = updated
// 			return models.Tasks[i], nil
// 		}
// 	}
// 	return models.Task{}, errors.New("task not found")
// }

//	func DeleteTask(id string) error {
//		for i, t := range models.Tasks {
//			if t.ID == id {
//				models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
//				return nil
//			}
//		}
//		return errors.New("task not found")
//	}
package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.IndentedJSON(http.StatusOK, tasks)
}
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}

func AddNewTask(c *gin.Context) {

	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	task := data.AddNewTask(newTask)
	c.IndentedJSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// for i, a := range models.Tasks {
	// 	if a.ID == id {
	// 		if updatedTask.Title != "" {
	// 			models.Tasks[i].Title = updatedTask.Title
	// 		}
	// 		if updatedTask.Description != "" {
	// 			models.Tasks[i].Description = updatedTask.Description
	// 		}
	// 		if updatedTask.Deadline != "" {
	// 			models.Tasks[i].Deadline = updatedTask.Deadline
	// 		}
	// 		// models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]... )
	// 		c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Updated"})
	// 		return
	// 	}
	// }
	_, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task Removed"})

}
