package models

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
}

var Tasks = []Task{
	{ID: "1", Title: "Eat dinner", Description: "nothing really", Deadline: "Nov 30"},
	{ID: "2", Title: "sleep well", Description: "at 5 pm", Deadline: "always"},
	{ID: "3", Title: "journal", Description: "always", Deadline: ""},
}
