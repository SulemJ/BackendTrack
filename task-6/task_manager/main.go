package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	router.StarterH()
	data.ConnectDB()
}
