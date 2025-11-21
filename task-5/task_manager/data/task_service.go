package data

import (
	"errors"
	"task_manager/models"
)

func GetAllTasks() []models.Task {
	return models.Tasks
}

func GetTaskById(id string) (models.Task, error) {
	for _, t := range models.Tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func AddNewTask(task models.Task) models.Task {
	models.Tasks = append(models.Tasks, task)
	return task
}

func UpdateTask(id string, updated models.Task) (models.Task, error) {
	for i, t := range models.Tasks {
		if t.ID == id {
			models.Tasks[i] = updated
			return models.Tasks[i], nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, t := range models.Tasks {
		if t.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
