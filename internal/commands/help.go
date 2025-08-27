package commands

import "fmt"

func ShowHelp() {
	helpText := 
	`
	Available commands:
	help        Show this help message
	add task    Add a task to the database (ticket, title, description)
	start task  Start a task. Prompts for task ID.
	exit        Exit the application
	`

	fmt.Println(helpText)
}