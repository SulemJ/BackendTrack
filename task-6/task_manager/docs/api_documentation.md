## A simple Task Management REST API using Go programming language and Gin Framework

## endPoints

# Get the full documentation of the API endpoints in here https://documenter.getpostman.com/view/34577456/2sB3dHVCUm

    - GET /tasks: Get a list of all tasks.
    - GET /tasks/:id: Get the details of a specific - task.
    - PUT /tasks/:id: Update a specific task. This  endpoint should accept a JSON body with the new details of the task.
    - DELETE /tasks/:id: Delete a specific task.
    - POST /tasks: Create a new task. This endpoint should accept a JSON body with the task's title, description, due date, and status.

## Folder Structure

    task_manager/
    ├── main.go
    ├── controllers/
    │   └── task_controller.go
    ├── models/
    │   └── task.go
    ├── data/
    │   └── task_service.go
    ├── router/
    │   └── router.go
    ├── docs/
    │   └── api_documentation.md
    └── go.mod

## MongoDB Integration Guide (Go + MongoDB Driver)

# 1. Prerequisites

    - Before running the project, ensure that:
    - Go ≥ 1.20 is installed
    - MongoDB database is running locally or on Atlas

# 2. Install Dependencies

    - go get go.mongodb.org/mongo-driver/mongo
    - go get go.mongodb.org/mongo-driver/mongo/options

    - Run go mod tidy to sync dependencies.

# 3. Configure MongoDB Connection

    - Create/load a connection string:

.env or environment variable - MONGO_URI=mongodb://localhost:27017 - DB_NAME=taskManager

OR if using Atlas:

    - mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority

# 4. Connect to MongoDB in Go

    import (
        "context"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
        "log"
    )
