package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"database/sql"
	"time"
	"github.com/manifoldco/promptui"
	"example.com/simple_cli/internal/db"
	"example.com/simple_cli/models"
	"example.com/simple_cli/internal/prompts"
)

func ShowAvailableTasks() models.Task {
    rows, err := db.DB.Query("SELECT ticket, title, description FROM tasks")
    if err != nil {
        fmt.Printf("Failed to retrieve tasks: %v\n", err)
        return models.Task{}
    }
    defer rows.Close()

    var tasks []models.Task

    // Read all rows into slice
    for rows.Next() {
        var t models.Task
        if err := rows.Scan(&t.Ticket, &t.Title, &t.Description); err != nil {
            fmt.Printf("Failed to scan task: %v\n", err)
            return models.Task{}
        }
        tasks = append(tasks, t)
    }

    if len(tasks) == 0 {
        fmt.Println("No tasks found")
        return models.Task{}
    }

    // Call selection function, passing the tasks slice
    return SelectAvailableTasks(tasks)
}

func SelectAvailableTasks(tasks []models.Task) models.Task {
    // Prepare display strings
    items := make([]string, len(tasks))
    for i, t := range tasks {
        items[i] = fmt.Sprintf("Ticket %d: %s", t.Ticket, t.Title)
    }

    // Create prompt
    prompt := promptui.Select{
        Label: "Select a task",
        Items: items,
    }

    index, _, err := prompt.Run()
    if err != nil {
        fmt.Printf("Prompt failed: %v\n", err)
        return models.Task{}
    }

    selectedTask := tasks[index]
    fmt.Printf("You selected Ticket %d: %s\n", selectedTask.Ticket, selectedTask.Title)

	return selectedTask
}

func ShowCurrentTask(rt **models.Task) {
	if *rt == nil {
		fmt.Println("No tasks are currently running.")
		return
	}

	fmt.Println("Currently Working On:")
	fmt.Printf("Ticket: %d | Title: %s | Description: %s\n", (*rt).Ticket, (*rt).Title, (*rt).Description)
}

func AddTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ticket number: ")
	ticketStr, _ := reader.ReadString('\n')
	ticketStr = strings.TrimSpace(ticketStr)

	var ticket int
	_, err := fmt.Sscanf(ticketStr, "%d", &ticket)
	if err != nil {
		fmt.Println("Invalid ticket number.")
		return
	}

	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	if title == "" || description == "" {
		fmt.Println("Title and description cannot be empty.")
		return
	}

	_, err = db.DB.Exec("INSERT INTO tasks (ticket, title, description) VALUES (?, ?, ?)", ticket, title, description)
	if err != nil {
		fmt.Printf("Failed to add task: %v\n", err)
		return
	}

	fmt.Printf("Task '%s' added successfully!\n", title)
}

func StartTask(rt **models.Task) {
    selected := ShowAvailableTasks()
    if selected.Ticket == 0 {
        return
    }

	var exists int
	err := db.DB.QueryRow("SELECT ticket FROM tasks WHERE ticket = ?", selected.Ticket).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No ticket found with ID %d.\n", selected.Ticket)
		} else {
			fmt.Printf("Failed to find ticket: %v\n", err)
		}
		return
	}

	if *rt != nil  {
		fmt.Printf("Task with Ticket %d is already running.\n", selected.Ticket)
		return
	}

	*rt = &models.Task{Ticket: selected.Ticket, StartTime: time.Now()}


	// Here you would typically update the task status in the database
	// For simplicity, we'll just print a message
	fmt.Printf("Task with ID %d started.\n", selected.Ticket)
}

func EndTask(rt **models.Task)bool {
	if *rt == nil {
		fmt.Println("No task is currently running.")
		return false
	}

	if prompts.YesNoPrompt("Are you sure you want to end the current task?") {
		(*rt) = nil
	}

	return (*rt) == nil
}