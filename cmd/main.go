package main

import (
    "fmt"
	"bufio"
	"os"
	"strings"
	"log"
	"time"
	"path/filepath"
	
    "example.com/simple_cli/internal/db"
    "example.com/simple_cli/internal/commands"
	"example.com/simple_cli/models"
)

var runningTask *models.Task

var date = time.Now().Format("2006-01-02")
var minutes = [4] string {"00", "15", "30", "45"}
var hours = [15] string {"08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22"}



func main() {
	initDb()

    reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to My CLI! Type 'help' for commands, 'exit' to quit.")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Available commands:")
	commands.ShowHelp()

	for {
		fmt.Print("> ") // prompt
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // remove newline


		switch input {
		case "--help":
			commands.ShowHelp()
		case "add task":
			commands.AddTask()
		case "current task":
			commands.ShowCurrentTask(&runningTask)
		case "start task":
			commands.StartTask(&runningTask)
		case "end task":
			ended := commands.EndTask(&runningTask)
			if ended {
				fmt.Println("Task successfully ended.")
			} else {
				fmt.Println("Task still running.")
			}
			reader.ReadString('\n') 
		default:
			fmt.Println("Unknown command... Type --help to get a list of possible commands", input)
		}
	}
}

func initDb() {
	dbPath := "./data/app.db"

		// Make sure parent folder exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create DB directory: %v", err)
	}

	err := db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
}
