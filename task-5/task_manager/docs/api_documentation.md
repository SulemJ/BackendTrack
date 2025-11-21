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
